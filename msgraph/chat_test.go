package msgraph

import (
	"testing"
)

func TestChat(t *testing.T) {

	err := SendMessage()

	if err != nil {
		t.Error(err)
	}

}
