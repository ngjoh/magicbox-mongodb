package kitchen

import (
	"testing"
)

func TestGetStatus(t *testing.T) {

	k, err := GetStatus("magicbox")
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
