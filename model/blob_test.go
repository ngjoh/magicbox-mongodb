package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlobAddString(t *testing.T) {
	err := SetBlobString("niels", "niels")

	assert.Nil(t, err)
}

func TestBlobGetString(t *testing.T) {
	b, err := GetBlob("niels")

	assert.Nil(t, err)
	assert.Equal(t, "niels", *b)
}
