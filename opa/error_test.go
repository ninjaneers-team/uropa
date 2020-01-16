package opa

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNotFoundErr(T *testing.T) {

	assert := assert.New(T)
	var e err404
	assert.True(IsNotFoundErr(e))
	assert.False(IsNotFoundErr(nil))

	err := errors.New("not a 404")
	assert.False(IsNotFoundErr(err))
}
