package file

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
	"github.com/stretchr/testify/assert"
)

func Test_ensureJSON(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			"empty array is kept as is",
			args{map[string]interface{}{
				"foo": []interface{}{},
			}},
			map[string]interface{}{
				"foo": []interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ensureJSON(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ensureJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadOpaStateFromStdinFailsToParseText(t *testing.T) {
	var filename = "-"
	assert := assert.New(t)
	assert.Equal("-", filename)

	var content bytes.Buffer
	content.Write([]byte("hunter2\n"))

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content.Bytes()); err != nil {
		panic(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		panic(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	c, err := GetContentFromFile(filename)
	assert.NotNil(err)
	assert.Nil(c)
}

func TestReadOpaStateFromStdin(t *testing.T) {
	var filename = "-"
	assert := assert.New(t)
	assert.Equal("-", filename)

	var content bytes.Buffer
	content.Write([]byte("policies:\n- raw: test.com\n  id: test policy\n"))

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content.Bytes()); err != nil {
		panic(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		panic(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	c, err := GetContentFromFile(filename)
	assert.NotNil(c)
	assert.Nil(err)

	assert.Equal(opa.Policy{
		ID:  opa.String("test policy"),
		Raw: opa.String("test.com"),
	},
		c.Policies[0].Policy)
}
