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

type AuthorHandler struct {
	services *services.AuthorService
}

func (a *AuthorHandler) CreateAuthor(r *gin.Context) {
	user_id := r.MustGet("user_id").(uint)
	author_input := dto.AuthorInputdto{}
	author_domain := domain.AuthorInput{}
	result := validator.Validate(r, &author_input, &author_domain)
	if result == nil {
		return
	}
	data, err := a.services.CreateAuthor(result, user_id)
	if err != nil {
		apperrors.ErrorFormat(r, err)
		return
	}
	r.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Authore registered",
		"data":    data,
	})

}

func (a *AuthorHandler) Initialize(r *gin.Engine) {
	h := r.Group("/api/v1/author")
	h.POST("", middleware.Authmiddleware(), a.CreateAuthor)

}

func NewAuthorHandler(r *gin.Engine, s *services.AuthorService) {
	h := &AuthorHandler{
		services: s,
	}
	h.Initialize(r)
}
