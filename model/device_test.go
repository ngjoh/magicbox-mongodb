package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncDevices(t *testing.T) {
	err := SyncDevices()

	assert.Nil(t, err)

}
