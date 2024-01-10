package kitchen

import (
	"testing"
)

func TestShip(t *testing.T) {
	output, err := Ship("meeting-infrastructure")
	if err != nil {
		t.Error(err)
	}
	t.Log(output)
}
