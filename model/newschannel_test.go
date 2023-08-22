package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsChannel(t *testing.T) {
	err := ImportNewsChannels()

	assert.Nil(t, err)
}
