package msgraph

import (
	"context"

	graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/spf13/viper"
)

func SendMessage() error {
	graphClient, err := ConnectUsingClientSecret(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"))
	if err != nil {
		return err
	}

	requestBody := graphmodels.NewChat()
	chatType := graphmodels.ONEONONE_CHATTYPE
	requestBody.SetChatType(&chatType)

	conversationMember := graphmodels.NewAadUserConversationMember()
	roles := []string{
		"owner",
	}
	conversationMember.SetRoles(roles)
	additionalData := map[string]interface{}{
		"odataBind": "https://graph.microsoft.com/v1.0/users('niels.johansen@nexigroup.com')",
	}
	conversationMember.SetAdditionalData(additionalData)
	conversationMember1 := graphmodels.NewAadUserConversationMember()
	roles = []string{
		"owner",
	}
	conversationMember1.SetRoles(roles)
	// additionalData := map[string]interface{}{
	//   "odataBind" : "https://graph.microsoft.com/v1.0/users('alex@contoso.com')",
	// }
	// conversationMember1.SetAdditionalData(additionalData)

	members := []graphmodels.ConversationMemberable{
		conversationMember,
		conversationMember1,
	}
	requestBody.SetMembers(members)

	_, err = graphClient.Chats().Post(context.Background(), requestBody, nil)

	return err
}
