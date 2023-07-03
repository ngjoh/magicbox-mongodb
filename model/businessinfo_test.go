package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCountry(t *testing.T) {
	err := NewCountry("niels", "niels")

	assert.Nil(t, err)

}

func TestGetCountries(t *testing.T) {
	c, err := Countries()

	assert.Nil(t, err)

	assert.Greater(t, len(c), 0)

}

func TestCreateUnit(t *testing.T) {
	err := NewBusinessGroupUnit("niels", "niels")

	assert.Nil(t, err)

}
