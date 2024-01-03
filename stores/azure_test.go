package stores

import (
	"testing"
)

func TestGitStatus(t *testing.T) {

	k, err := AzureBlobStorages()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
