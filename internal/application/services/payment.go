package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	payload "github.com/marveldo/gogin/internal/application/payloads"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/application/utils"
	"github.com/marveldo/gogin/internal/db"
)

type PaymentService struct {
	R  *repository.PaymentRepository
	U  *repository.Userrespository
	C  *repository.CartRepository
	AC *asynq.Client
}

func (p *PaymentService) PlaceOrder(user_id uint, cart_domain *domain.CreatePaymentDomain, g *gin.Context) (*dto.PaymentCreatedResponseDto, error) {
	user, err := p.U.GetUser(&domain.GetUserQuery{ID: &user_id})
	if err != nil {
		return nil, err
	}
	if user.Cart.ID != cart_domain.CartID {
		return nil, errors.New("Wrong CartId Submitted")
	}
	cart_items, err := p.C.GetCartItems(user.Cart.ID)
	if err != nil {
		return nil, err
	}
	order_item, err, tx, total := p.R.CreatePaymentOrder(cart_items, user_id)

	if err != nil {
		return nil, err
	}
	defer func() {
		if tx != nil {
			tx.Rollback()
		}
	}()
	if total <= 0 {
		return nil, errors.New("Total Must be greater than zero")
	}

	resp, err := utils.CallPaystackUrl(g, user.Email, total, "", order_item.Reference)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resp_body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var paystackResp map[string]interface{}
	if err := json.Unmarshal(resp_body_bytes, &paystackResp); err != nil {
		return nil, err
	}
	authURL, ok := paystackResp["data"].(map[string]interface{})["authorization_url"].(string)
	if !ok {
		return nil, errors.New("Something Went Wrong With the Response")
	}

	order_response := dto.OrderDto{}
	err = copier.Copy(&order_response, order_item)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	tx = nil
	return &dto.PaymentCreatedResponseDto{
		Order:   order_response,
		AuthUrl: authURL,
	}, nil
}

func (p *PaymentService) UpdateOrder(payment_domain *domain.PaymentWebhookdomain) (*db.OrderModel, error) {
	var books []db.Bookmodel
	order_item, tx, err := p.R.UpdatePaymentOrder((payment_domain))
	if err != nil {
		return nil, err
	}
	defer func() {
		if tx != nil {
			tx.Rollback()
		}
	}()
	switch order_item.Status {
	case db.Success:
		user, err := p.U.GetUserWithtx(tx, &domain.GetUserQuery{ID: &order_item.UserId})
		if err != nil {
			return nil, err
		}
		cart_items, err := p.C.GetCartItemsWithtx(tx, user.Cart.ID)
		if err != nil {
			return nil, err
		}
		for _, cart := range cart_items {
			books = append(books, cart.Book)
		}
		err = tx.Model(user).Association("Books").Append(books)
		if err != nil {
			return nil, err
		}
		cart := user.Cart
        
		if err := tx.Where("cart_id = ?", cart.ID).Delete(&cart.CartItems).Error; err != nil {
			return nil, err
		}

		email_payload := payload.EmailPayload{
			Email:    user.Email,
			Username: user.Username,
		}
		b, err := json.Marshal(email_payload)
		if err != nil {
			return nil, err
		}
		task := asynq.NewTask("success-email", b)
		p.AC.Enqueue(task)
		tx.Commit()
		tx = nil
		return order_item, nil
	case db.Failed:
		tx.Commit()
		tx = nil
		return order_item, nil

	default:
		tx.Commit()
		tx = nil
		return order_item, nil
	}

}
