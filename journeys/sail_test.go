package journeys

import (
	"testing"
)

func TestSail(t *testing.T) {
	err := Sail()
	if err != nil {
		t.Error(err)
	}

}
