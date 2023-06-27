package msgraph

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestDeviceCodeAuth(t *testing.T) {

	client, err := ConnectUsingDeviceCode(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"))

	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, client)

}

func TestClientSecreAuth(t *testing.T) {

	client, err := ConnectUsingClientSecret(viper.GetString("SPAUTH_DOMAIN"), viper.GetString("SPAUTH_CLIENTID"), viper.GetString("SPAUTH_CLIENTSECRET"))

	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, client)

}
