package services

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/application/utils"
)

type PaymentService struct {
	R *repository.PaymentRepository
	U *repository.Userrespository
	C *repository.CartRepository
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
	resp, err := utils.CallPaystackUrl(g, user.Email, total, "", order_item.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer resp.Body.Close()

	resp_body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	var paystackResp map[string]interface{}
	if err := json.Unmarshal(resp_body_bytes, &paystackResp); err != nil {
		tx.Rollback()
		return nil, err
	}
	authURL, ok := paystackResp["data"].(map[string]interface{})["authorization_url"].(string)
	if !ok {
		tx.Rollback()
		return nil, errors.New("Something Went Wrong With the Response")
	}

	order_response := dto.OrderDto{}
	err = copier.Copy(&order_response, order_item)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &dto.PaymentCreatedResponseDto{
		Order:   order_response,
		AuthUrl: authURL,
	}, nil
}
