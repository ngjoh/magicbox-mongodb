package stores

import (
	"testing"
)

func TestMongoClusters(t *testing.T) {

	k, err := PerconaCRDS()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
