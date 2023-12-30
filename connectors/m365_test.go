package connectors

import (
	"testing"
)

func TestM365Context(t *testing.T) {

	k, err := M365Context()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}

func TestM365Sites(t *testing.T) {

	k, err := M365Sites()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
