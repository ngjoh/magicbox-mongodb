package connectors

import (
	"testing"
)

func TestSharePointTenants(t *testing.T) {

	k, err := SharePointTenants()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
