package model

import (
	//"log"

	"testing"
)

func TestCreateUser(t *testing.T) {

	t.Log("Syncing Site Navigation")
	err := SyncSitesNavigation()
	if err != nil {
		t.Error(err)
	}
	t.Log("Done")
}
