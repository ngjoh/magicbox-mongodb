package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompanies(t *testing.T) {
	err := ImportCompanies()

	assert.Nil(t, err)
}
