package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	state, err := NewOpaState()
	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(state)
}

func state() *OpaState {
	s, err := NewOpaState()
	if err != nil {
		panic(err)
	}
	return s
}
