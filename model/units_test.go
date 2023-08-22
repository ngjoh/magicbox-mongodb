package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnits(t *testing.T) {
	err := ImportUnits()

	assert.Nil(t, err)
}
