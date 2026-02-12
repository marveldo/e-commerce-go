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

type WaitlistHandler struct {
	services *services.WaitlistService
}

func (h *WaitlistHandler) GetBooksInWaitlist(g *gin.Context) {
	userId := g.MustGet("user_id").(uint)
	books, err := h.services.GetBooksInWaitlist(userId)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Waitlist Retrieved Succesfully",
		"data":    books,
	})
}

func (h *WaitlistHandler) AddBookToWaitlist(g *gin.Context) {
	userId := g.MustGet("user_id").(uint)
	input := domain.AddBooktoWaitlist{}
	input_dto := dto.AddBooktoWaitlist{}
	result := validator.Validate(g, &input_dto, &input)
	if result == nil {
		return
	}
	book, err := h.services.AddBookToWaitlist(userId, result.Book_id)
	if err != nil {
		apperrors.ErrorFormat(g, err)
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Book Added to Waitlist Succesfully",
		"data":    book,
	})
}

func (h *WaitlistHandler) Initialize(r *gin.Engine) {
	waitlist := r.Group("/api/v1/waitlist")
	waitlist.GET("/", middleware.Authmiddleware(), h.GetBooksInWaitlist)
	waitlist.POST("/add", middleware.Authmiddleware(), h.AddBookToWaitlist)

}

func NewWaitlistHandler(r *gin.Engine, s *services.WaitlistService) {
	h := &WaitlistHandler{
		services: s,
	}
	h.Initialize(r)
}
