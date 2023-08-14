package msgraph

import (
	"fmt"

	"context"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

func ConnectUsingDeviceCode(tenantId string, clientId string) (*msgraphsdk.GraphServiceClient, error) {

	cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID: tenantId,
		ClientID: clientId,
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(fmt.Sprint(message.Message))

			return nil
		},
	})

	if err != nil {
		return nil, err

	}
	return msgraphsdk.NewGraphServiceClientWithCredentials(cred, nil)

}
func ConnectUsingClientSecret(tenantId string, clientId string, clientSecret string) (*msgraphsdk.GraphServiceClient, error) {

	cred, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, &azidentity.ClientSecretCredentialOptions{})

	if err != nil {
		return nil, err

	}
	return msgraphsdk.NewGraphServiceClientWithCredentials(cred, nil)

}
func printOdataError(err error) {
	switch err.(type) {
	case *odataerrors.ODataError:
		typed := err.(*odataerrors.ODataError)
		fmt.Println("error:", typed.Error())
		if terr := typed.Error(); terr != "" {

			fmt.Printf("msg: %s", terr)
		}
	default:
		fmt.Printf("%T > error: %#v", err, err)
	}
}
