package msgraph

import (
	"testing"

	"github.com/spf13/viper"
)

func TestMail(t *testing.T) {
	client, err := ConnectUsingClientSecret(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"))
	s := "test"
	err = SendMailMessage(client, &s, &s, &s)

	if err != nil {
		t.Error(err)
	}

}
