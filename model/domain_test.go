package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncDomains(t *testing.T) {
	err := SyncDomains()

	assert.Nil(t, err)

}
