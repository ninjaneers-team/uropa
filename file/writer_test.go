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
			Workspace:  "foo",
			Filename:   filename,
			FileFormat: YAML,
		})
	})
	assert.Equal("_format_version: \"1.1\"\n_workspace: foo\n", output)
	// JSON
	output = captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Workspace:  "foo",
			Filename:   filename,
			FileFormat: JSON,
		})
	})
	expected := `{
  "_format_version": "1.1",
  "_workspace": "foo"
}`
	assert.Equal(expected, output)
}

func TestWriteOpaStateToStdoutStateWithOneService(t *testing.T) {
	var ks, _ = state.NewOpaState()
	var filename = "-"
	assert := assert.New(t)
	var service state.Service
	service.ID = Opa.String("first")
	service.Host = Opa.String("example.com")
	service.Name = Opa.String("my-policy")
	ks.Services.Add(service)
	// YAML
	output := captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Filename:   filename,
			FileFormat: YAML,
		})
	})
	expected := fmt.Sprintf("_format_version: \"1.1\"\nservices:\n- host: %s\n  name: %s\n", *service.Host, *service.Name)
	assert.Equal(expected, output)
	// JSON
	output = captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Workspace:  "foo",
			Filename:   filename,
			FileFormat: JSON,
		})
	})
	expected = `{
  "_format_version": "1.1",
  "_workspace": "foo",
  "services": [
    {
      "host": "example.com",
      "name": "my-policy"
    }
  ]
}`
	assert.Equal(expected, output)
}

func TestWriteOpaStateToStdoutStateWithOneServiceOneRoute(t *testing.T) {
	var ks, _ = state.NewOpaState()
	var filename = "-"
	assert := assert.New(t)
	var service state.Service
	service.ID = Opa.String("first")
	service.Host = Opa.String("example.com")
	service.Name = Opa.String("my-policy")
	ks.Services.Add(service)

	var route state.Route
	route.Name = Opa.String("my-route")
	route.ID = Opa.String("first")
	route.Hosts = Opa.StringSlice("example.com", "demo.example.com")
	route.Service = &Opa.Service{
		ID:   Opa.String(*service.ID),
		Name: Opa.String(*service.Name),
	}

	ks.Routes.Add(route)
	// YAML
	output := captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Filename:   filename,
			FileFormat: YAML,
		})
	})
	expected := fmt.Sprintf(`_format_version: "1.1"
services:
- host: %s
  name: %s
  routes:
  - hosts:
    - %s
    - %s
    name: %s
`, *service.Host, *service.Name, *route.Hosts[0], *route.Hosts[1], *route.Name)
	assert.Equal(expected, output)
	// JSON
	output = captureOutput(func() {
		OpaStateToFile(ks, WriteConfig{
			Workspace:  "foo",
			Filename:   filename,
			FileFormat: JSON,
		})
	})
	expected = `{
  "_format_version": "1.1",
  "_workspace": "foo",
  "services": [
    {
      "host": "example.com",
      "name": "my-policy",
      "routes": [
        {
          "hosts": [
            "example.com",
            "demo.example.com"
          ],
          "name": "my-route"
        }
      ]
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
