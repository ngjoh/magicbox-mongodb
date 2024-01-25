package kitchen

import (
	"testing"
)

func TestGetStatus(t *testing.T) {

	k, err := GetStatus("magicbox", true)
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
