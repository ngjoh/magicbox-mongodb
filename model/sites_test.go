package model

import (
	//"log"

	"testing"
)

func TestSyncHubSiteNavigation(t *testing.T) {

	t.Log("Syncing Site Navigation")
	err := SyncSitesNavigation()
	if err != nil {
		t.Error(err)
	}
	t.Log("Done")
}

func TestSyncHubSitePages(t *testing.T) {
	t.Log("Syncing Site Pages")

	err := SyncHubSitePages("b80f09f2-c5e5-4f69-9944-33e8fe18a96c")
	if err != nil {
		t.Error(err)
	}

	t.Log("Done")
}
