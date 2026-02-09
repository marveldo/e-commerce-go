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

type Userhandler struct {
	services *services.Userservice
}

func (u *Userhandler) Register(r *gin.Context) {
	input := domain.UserInput{}
	input_dto := dto.UserInputDto{}
	result := validator.Validate(r, &input_dto, &input)
	if result == nil {
		return
	}
	user, err := u.services.Create(result)
	if err != nil {
		apperrors.ErrorFormat(r, err)
		return
	}
	r.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "User Created Successfully",
		"data":    user,
	})
}

func (u *Userhandler) Login(r *gin.Context) {
	input := domain.LoginInput{}
	input_dto := dto.LoginInputDto{}
	result := validator.Validate(r, &input_dto, &input)
	if result == nil {
		return
	}
	data, err := u.services.AuthenticateUser(result)
	if err != nil {
		apperrors.ErrorFormat(r, err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User Login Successful",
		"data":    data,
	})
}

func (u *Userhandler) GetLoginUser(r *gin.Context) {
	username := r.MustGet("username").(string)
	query := &domain.GetUserQuery{
		Username: &username,
	}
	user, err := u.services.GetUser(query)
	if err != nil {
		apperrors.ErrorFormat(r, err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User Retrieved",
		"data":    user,
	})
}

func (u *Userhandler) GoogleLogin(r *gin.Context) {
	input := dto.GoogleLoginInputDto{}
	input_domain := domain.GoogleLoginDomain{}
	result := validator.Validate(r, &input, &input_domain)
	if result == nil {
		return
	}
	data, err := u.services.GoogleLogin(result)
	if err != nil {
		apperrors.ErrorFormat(r, err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User Login Successful",
		"data":    data,
	})
}
func (u *Userhandler) Initialize(r *gin.Engine) {
	hg := r.Group("/api/v1")
	hg.POST("/user", u.Register)
	hg.POST("/login", u.Login)
	hg.GET("/user", middleware.Authmiddleware(), u.GetLoginUser)
	hg.POST("/google", u.GoogleLogin)

}

func NewUserHandler(r *gin.Engine, s *services.Userservice) {
	h := &Userhandler{
		services: s,
	}
	h.Initialize(r)
}
