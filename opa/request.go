package opa

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// NewFile creates a request based on the inputs.
// endpoint should be relative to the baseURL specified during
// client creation.
// body is always marshaled into JSON.
func (c *Client) NewFile(body string) (*os.File, error) {
	f, err := ioutil.TempFile(os.TempDir(), "uropa-*.rego")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	if _, err = f.WriteString(body); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	return f, nil
}

// NewFormEncodedRequest creates a request based on the inputs.
// endpoint should be relative to the baseURL specified during
// client creation.
func (c *Client) NewFormEncodedRequest(method, endpoint string, qs interface{},
	reader io.Reader) (*http.Request, error) {

	if endpoint == "" {
		return nil, errors.New("endpoint can't be nil")
	}
	//Create a new request
	req, err := http.NewRequest(method, c.baseURL+endpoint,
		reader)
	if err != nil {
		return nil, err
	}

	// add body if needed
	//if reader != nil {
	//	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//}

	// add query string if any
	if qs != nil {
		values, err := query.Values(qs)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = values.Encode()
	}
	return req, nil
}

// NewJSONRequest creates a request based on the inputs.
// endpoint should be relative to the baseURL specified during
// client creation.
// body is always marshaled into JSON.
func (c *Client) NewJSONRequest(method, endpoint string, qs interface{},
	body interface{}) (*http.Request, error) {

	if endpoint == "" {
		return nil, errors.New("endpoint can't be nil")
	}
	//body to be sent in JSON
	var buf []byte
	if body != nil {
		var err error
		buf, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	//Create a new request
	req, err := http.NewRequest(method, c.baseURL+endpoint,
		bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	// add body if needed
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	// add query string if any
	if qs != nil {
		values, err := query.Values(qs)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = values.Encode()
	}
	return req, nil
}
