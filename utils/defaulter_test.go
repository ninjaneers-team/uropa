package utils

import (
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
	"github.com/stretchr/testify/assert"
)

func TestDefaulter(t *testing.T) {
	assert := assert.New(t)

	var d Defaulter

	assert.NotNil(d.Register(nil))
	assert.NotNil(d.Set(nil))

	assert.Panics(func() {
		d.MustSet(d)
	})

	type Foo struct {
		A string
		B []string
	}
	defaultFoo := &Foo{
		A: "defaultA",
		B: []string{"default1"},
	}
	assert.Nil(d.Register(defaultFoo))

	// sets a default
	var arg Foo
	assert.Nil(d.Set(&arg))
	assert.Equal("defaultA", arg.A)
	assert.Equal([]string{"default1"}, arg.B)

	// doesn't set a default
	arg1 := Foo{
		A: "non-default-value",
	}
	assert.Nil(d.Set(&arg1))
	assert.Equal("non-default-value", arg1.A)

	// errors on an unregistered type
	type Bar struct {
		A string
	}
	assert.NotNil(d.Set(&Bar{}))
	assert.Panics(func() {
		d.MustSet(&Bar{})
	})
}

func TestPolicySetTest(t *testing.T) {
	assert := assert.New(t)
	d, err := GetOpaDefaulter()
	assert.NotNil(d)
	assert.Nil(err)

	testCases := []struct {
		desc string
		arg  *opa.Policy
		want *opa.Policy
	}{
		{
			desc: "empty service",
			arg:  &opa.Policy{},
			want: &policyDefaults,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := d.Set(tC.arg)
			assert.Nil(err)
			assert.Equal(tC.want, tC.arg)
		})
	}
}
