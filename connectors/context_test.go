package connectors

import (
	"testing"
)

func TestGetContext(t *testing.T) {

	k, err := GetContext()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
