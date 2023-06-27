package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS(t *testing.T) {
	err := NewSubscription("test", "https://christianiabpos.sharepoint.com/sites/Cava3", "Test Changes")

	assert.Nil(t, err)
}
