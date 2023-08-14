package msgraph

import (
	"context"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func SendMailMessage(client *msgraphsdk.GraphServiceClient, subject string, body string, recipient string) error {

	message := models.NewMessage()
	message.SetSubject(&subject)

	messageBody := models.NewItemBody()
	messageBody.SetContent(&body)
	contentType := models.TEXT_BODYTYPE
	messageBody.SetContentType(&contentType)
	message.SetBody(messageBody)

	toRecipient := models.NewRecipient()
	emailAddress := models.NewEmailAddress()
	emailAddress.SetAddress(&recipient)
	toRecipient.SetEmailAddress(emailAddress)
	message.SetToRecipients([]models.Recipientable{
		toRecipient,
	})

	sendMailBody := users.NewItemSendMailPostRequestBody()
	sendMailBody.SetMessage(message)

	// mailTips, err := c.Me().GetMailTips().Post(context.Background(), nil, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// log.Println(mailTips)
	client.Domains().Get(context.Background(), nil)
	return client.Me().SendMail().Post(context.Background(), sendMailBody, nil)

}

func GetDomains(client *msgraphsdk.GraphServiceClient) (models.DomainCollectionResponseable, error) {

	return client.Domains().Get(context.Background(), nil)

}
