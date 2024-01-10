package connectors

import (
	"testing"
)

func TestCentrifugo(t *testing.T) {

	err := CentrifugePost("echo11", "12314", false)
	if err != nil {
		t.Error(err)
	}

}
