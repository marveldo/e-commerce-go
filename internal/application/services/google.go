package services

import (
	"context"

	"cloud.google.com/go/auth/credentials/idtoken"
	"github.com/marveldo/gogin/internal/config"
)

func VerifyOauthGoogleToken(token string) (*idtoken.Payload , error) {
	ctx := context.Background()
	payload , err := idtoken.Validate(ctx , token , config.LoadConfig().Google_Client_Id)
	if err != nil {
		return nil , err
	}
	return payload , nil

}
