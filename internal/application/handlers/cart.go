package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	apperrors "github.com/marveldo/gogin/internal/application/errors"
	"github.com/marveldo/gogin/internal/application/middleware"
	"github.com/marveldo/gogin/internal/application/services"
	"github.com/marveldo/gogin/internal/application/validator"
)

type CartHandler struct {
	s *services.CartService
}

func (c *CartHandler) GetCartItems(g *gin.Context) {
	user_id := g.MustGet("user_id").(uint)
	items, err := c.s.GetCartItems(user_id)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(200, gin.H{
		"code": 200,
		"data": items,
	})
}

func (c *CartHandler) AddCartItem(g *gin.Context) {
	user_id := g.MustGet("user_id").(uint)
	input_domain := &domain.CartItemInputDomain{}
	cart_input := &dto.CartItemInputDto{}
	result := validator.Validate(g, cart_input, input_domain)
	if result == nil {
		return
	}
	item, err := c.s.AddCartItem(user_id, result)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(200, gin.H{
		"code": 200,
		"data": item,
	})

}

func (c *CartHandler) Initialize(r *gin.Engine) {
	cart := r.Group("api/v1/cart")
	cart.GET("", middleware.Authmiddleware(), c.GetCartItems)
	cart.POST("/add", middleware.Authmiddleware(), c.AddCartItem)
}

func NewCartHandler(r *gin.Engine, s *services.CartService) {
	handler := &CartHandler{
		s: s,
	}
	handler.Initialize(r)
}
