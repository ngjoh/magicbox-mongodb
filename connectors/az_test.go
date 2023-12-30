package connectors

import (
	"testing"
)

func TestAzureSubscriptions(t *testing.T) {

	k, err := AzureSubscriptions()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
