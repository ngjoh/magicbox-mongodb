package officegraph

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/spf13/viper"
)

func GetClient() (*ClientWithResponses, string, error) {
	//client, err := ConnectUsingClientSecret(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"))
	cred, err := azidentity.NewClientSecretCredential(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"), &azidentity.ClientSecretCredentialOptions{})
	if err != nil {
		return nil, "", err
	}
	ctx := context.Background()
	opts := &policy.TokenRequestOptions{Scopes: []string{"https://graph.microsoft.com/.default"}, TenantID: viper.GetString("SPAUTH_DOMAIN")}
	token, err := cred.GetToken(ctx, *opts)
	if err != nil {
		return nil, "", err
	}
	// Example BearerToken
	// See: https://swagger.io/docs/specification/authentication/bearer-authentication/
	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(token.Token)
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}

	// Exhaustive list of some defaults you can use to initialize a Client.
	// If you need to override the underlying httpClient, you can use the option
	//
	// WithHTTPClient(httpClient *http.Client)
	//

	client, _ := NewClient("https://graph.microsoft.com/v1.0/", WithRequestEditorFn(bearerTokenProvider.Intercept))
	c := &ClientWithResponses{client}
	return c, token.Token, nil

}
