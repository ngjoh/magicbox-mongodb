package msgraph

import (
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestMail(t *testing.T) {
	client, err := ConnectUsingClientSecret(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"))
	err = SendMailMessage(client, "Subject", "Body", "niels.johansen@nexigroup.com")

	if err != nil {
		t.Error(err)
	}

}

func TestGetDomains(t *testing.T) {
	client, err := ConnectUsingClientSecret(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"))
	domainsResult, err := GetDomains(client)

	if err != nil {
		t.Error(err)
	}

	domains := domainsResult.GetValue()

	for _, domain := range domains {
		log.Println(fmt.Sprintf("days %d", domain.GetPasswordNotificationWindowInDays()))
	}

}
