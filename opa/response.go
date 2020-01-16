package opa

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Response is a Opa Admin API response. It wraps http.Response.
type Response struct {
	*http.Response
	//other Opa specific fields
}

func newResponse(res *http.Response) *Response {
	return &Response{Response: res}
}

func hasError(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 399 {
		return nil
	}

	if res.StatusCode == 404 {
		return err404{}
	}

	body, _ := ioutil.ReadAll(res.Body) // TODO error in error?
	return errors.New(res.Status + " " + string(body))
}
