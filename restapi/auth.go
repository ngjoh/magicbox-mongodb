package restapi

import (
	"context"
	"errors"
	"time"

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
		return "", errors.New("you cannot get an access token with that app key")
	}
	t := time.Now()
	t.Add(time.Minute * 10)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app":         app,
		"permissions": permissions,
		"expire":      t.UTC(),
	})
	model.LogAudit(app, permissions)
	tokenString, err = token.SignedString(mySigningKey)
	return tokenString, nil
}
func signin() usecase.Interactor {
	type SigninRequest struct {
		AppKey string `json:"appkey" binding:"required"`
	}
	type SigninResponse struct {
		AccessToken string `json:"token" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SigninRequest, output *SigninResponse) error {

		tokenString, err := IssueAccessToken(input.AppKey)
		*&output.AccessToken = tokenString
		return err

	})

	u.SetTitle("Authenticating app")
	u.SetDescription(`

You need an ´appkey´  to access the API. The ´appkey´  is issue by [niels.johansen@nexigroup.com](mailto:niels.johansen@nexigroup.com).

Use the ´appkey´  to get an access token through the /v1/authorize end point. The access token is valid for 10 minutes.

Pass the access token in the Authorization header as a Bearer token to access the API.


`)
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(authenticationTag)
	return u
}
