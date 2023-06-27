package msgraph

import (
	"log"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func SendMailMessage(c *msgraphsdk.GraphServiceClient, subject *string, body *string, recipient *string) error {

	message := models.NewMessage()
	message.SetSubject(subject)

	messageBody := models.NewItemBody()
	messageBody.SetContent(body)
	contentType := models.TEXT_BODYTYPE
	messageBody.SetContentType(&contentType)
	message.SetBody(messageBody)

	toRecipient := models.NewRecipient()
	emailAddress := models.NewEmailAddress()
	emailAddress.SetAddress(recipient)
	toRecipient.SetEmailAddress(emailAddress)
	message.SetToRecipients([]models.Recipientable{
		toRecipient,
	})

	sendMailBody := users.NewItemSendMailPostRequestBody()
	sendMailBody.SetMessage(message)

	// Send the message
	u := c.Users().ByUserId("niels.johansen@nexigroup.com")

	mt := u.GetMailTips()
	log.Println(mt)
	//.userClient.Me().SendMail().Post(context.Background(), sendMailBody, nil)

	return nil
}
