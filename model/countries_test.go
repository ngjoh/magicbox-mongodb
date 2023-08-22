package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountries(t *testing.T) {
	err := ImportCountries()

	assert.Nil(t, err)
}
