package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getString() string {
	return "TEST_STRING"
}

func TestInit(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(getString(), "TEST_STRING")
	assert.NotEqual(getString(), "")
}
