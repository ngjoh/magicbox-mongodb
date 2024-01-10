package kitchen

import (
	"testing"
)

func TestGetMetadata(t *testing.T) {
	metadata, err := GetMetadata("sharepoint-branding", "install", "10 get-appcatalogueurl.ps1")
	if err != nil {
		t.Error(err)
	}
	t.Log(metadata)

}
