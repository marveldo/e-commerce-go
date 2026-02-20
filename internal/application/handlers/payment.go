package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	apperrors "github.com/marveldo/gogin/internal/application/errors"
	"github.com/marveldo/gogin/internal/application/middleware"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/application/validator"
)

type PaymentHandler struct {
	services *services.PaymentService
}

func (p *PaymentHandler) TriggerPayment(c *gin.Context) {
	user_id := c.MustGet("user_id").(uint)
	input_dto := dto.CreatePaymentDto{}
	input_domain := domain.CreatePaymentDomain{}
	result := validator.Validate(c, &input_dto, &input_domain)
	if result == nil {
		return
	}
	data, err := p.services.PlaceOrder(user_id, result, c)
	if err != nil {
		apperrors.ErrorFormat(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Order Successful",
		"data":    data,
	})
}

func (p *PaymentHandler) Paymenthook(c *gin.Context) {
	input_dto := dto.PaymentWebhookdto{}
	input_domain := domain.PaymentWebhookdomain{}
	result := validator.Validate(c, &input_dto, &input_domain)
	if result == nil {
		return
	}
	data, err := p.services.UpdateOrder(result)
	if err != nil {
		apperrors.ErrorFormat(c, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Order Successful",
		"data":    data,
	})

}

func (p *PaymentHandler) Initialize(r *gin.Engine) {
	h := r.Group("/api/v1/payment")
	allowed_cidrs := []string{"52.31.139.75", "52.49.173.169", "52.214.14.220"}
	h.POST("", middleware.Authmiddleware(), p.TriggerPayment)
	h.POST("/hook", middleware.IpwhitelistMiddleware(allowed_cidrs), p.Paymenthook)
}

func NewPaymentHandler(r *gin.Engine, s *services.PaymentService) {
	handler := &PaymentHandler{
		services: s,
	}
	handler.Initialize(r)
}
