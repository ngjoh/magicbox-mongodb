package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS(t *testing.T) {
	//err := NewSubscription("test", "https://christianiabpos.sharepoint.com/sites/Cava3", "Test Changes", "https://niels-mac.nets-intranets.com/api/v1/subscription/notify")
	err := NewSubscription("test", "https://christianiabpos.sharepoint.com/sites/Cava3", "Test Changes", "https://magicbox.nexi-intra.com/api/v1/subscription/notify")

	assert.Nil(t, err)
}
