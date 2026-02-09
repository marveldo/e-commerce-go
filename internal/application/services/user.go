package services

import (
	"errors"
    "github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/application/utils"
	"github.com/marveldo/gogin/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type Userservice struct {
	R *repository.Userrespository
}

func (u *Userservice) Create(i *domain.UserInput) (*db.UserModel, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(i.Password), 10)
	if err != nil {
		return nil, err
	}
	i.Password = string(hashed_password)
	return u.R.Createuser(i)
}

func (u *Userservice) AuthenticateUser(i *domain.LoginInput) (*domain.LoginResponse, error) {
	query := &domain.GetUserQuery{
		Email: &i.Email,
	}
	user, err := u.R.GetUser(query)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(i.Password))
	if err != nil {
		return nil, errors.New("Password Not Correct")
	}
	access_token, err := utils.GenrateJwtToken(user.Username, user.ID, string(utils.ACCESS))
	if err != nil {
		return nil, err
	}
	refresh_token, err := utils.GenrateJwtToken(user.Username, user.ID, string(utils.REFRESH))
	if err != nil {
		return nil, err
	}
	login_response := &domain.LoginResponse{User: user, Access: access_token, Refresh: refresh_token}
	return login_response, nil
}

func (u *Userservice) GetUser(d *domain.GetUserQuery) (*db.UserModel, error) {
	return u.R.GetUser(d)
}

func (u *Userservice) GoogleLogin(i *domain.GoogleLoginDomain) (*domain.LoginResponse, error) {
	user_info, err := VerifyOauthGoogleToken(i.IDtoken)
	if err != nil {
		return nil, err
	}
	email := user_info.Claims["email"].(string)

	query := &domain.GetUserQuery{
		Email: &email,
	}
	user, err := u.R.GetUser(query)
	if err != nil {
		return nil, err
	}
	access_token, err := utils.GenrateJwtToken(user.Username, user.ID, string(utils.ACCESS))
	if err != nil {
		return nil, err
	}
	refresh_token, err := utils.GenrateJwtToken(user.Username, user.ID, string(utils.REFRESH))
	if err != nil {
		return nil, err
	}
	login_response := &domain.LoginResponse{User: user, Access: access_token, Refresh: refresh_token}
	return login_response, nil
}
