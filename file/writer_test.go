package file

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = writer
	os.Stderr = writer

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestWriteOpaStateToStdoutEmptyState(t *testing.T) {
	var ks, _ = state.NewOpaState()
	var filename = "-"
	assert := assert.New(t)
	assert.Equal("-", filename)
	assert.NotEmpty(t, ks)
	// YAML
	output := captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Filename:   filename,
			FileFormat: YAML,
		})
	})
	assert.Equal("_format_version: \"1.1\"\n", output)
	// JSON
	output = captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Filename:   filename,
			FileFormat: JSON,
		})
	})
	expected := `{
  "_format_version": "1.1"
}`
	assert.Equal(expected, output)
}

func TestWriteOpaStateToStdoutStateWithOnePolicy(t *testing.T) {
	var ks, _ = state.NewOpaState()
	var filename = "-"
	assert := assert.New(t)
	var policy state.Policy
	policy.ID = opa.String("first")
	policy.Raw = opa.String("example.com")
	ks.Policies.Add(policy)
	// YAML
	output := captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Filename:   filename,
			FileFormat: YAML,
		})
	})
	expected := fmt.Sprintf("_format_version: \"1.1\"\npolicies:\n  - id: %s\n    raw: %s\n", *policy.ID, *policy.Raw)
	assert.Equal(expected, output)
	// JSON
	output = captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Filename:   filename,
			FileFormat: JSON,
		})
	})
	expected = `{
  "_format_version": "1.1",
  "policies": [
    {
      "id": "first",
      "raw": "example.com"
    }
  ]
}`
	assert.Equal(expected, output)
}

func Test_addExtToFilename(t *testing.T) {
	type args struct {
		filename string
		format   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				filename: "foo",
				format:   "yolo",
			},
			want: "foo.yolo",
		},
		{
			args: args{
				filename: "foo.json",
				format:   "yolo",
			},
			want: "foo.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addExtToFilename(tt.args.filename, tt.args.format); got != tt.want {
				t.Errorf("addExtToFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
