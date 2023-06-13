package restapi

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

var authenticationTag = "Authentication"

var mySigningKey = []byte("AllYourBase")

func IssueIdToken(appName string, appSecret string) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app": appName,
		"key": appSecret,
		"nbf": 234,
	})

	tokenString, err = token.SignedString(mySigningKey)

	return tokenString, err
}

func IssueAccessToken(idToken string) (tokenString string, err error) {

	app, secret, err := ParseIdToken(idToken)
	valid, permissions := model.Authenticate(app, secret)
	if !valid {
		return "", errors.New("you've got access")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app":         app,
		"permissions": permissions,
	})
	model.LogAudit(app, permissions)
	tokenString, err = token.SignedString(mySigningKey)
	return tokenString, nil
}
func signin() usecase.Interactor {
	type SigninRequest struct {
		IdToken string `json:"idtoken" binding:"required"`
	}
	type SigninResponse struct {
		AccessToken string `json:"token" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SigninRequest, output *SigninResponse) error {

		tokenString, err := IssueAccessToken(input.IdToken)
		*&output.AccessToken = tokenString
		return err

	})

	u.SetTitle("Authenticating app")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(authenticationTag)
	return u
}
