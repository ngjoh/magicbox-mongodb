package connectors

import (
	"testing"
)

func TestGitStatus(t *testing.T) {

	k, err := GitStatus()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
