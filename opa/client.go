package opa

import (
	"context"
	"encoding/json"
	"github.com/ninjaneers-team/uropa/opa/custom"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const defaultBaseURL = "http://localhost:8001"

var pretty = false

type service struct {
	client *Client
}

var (
	defaultCtx = context.Background()
)

// Client talks to the Admin API or control plane of a
// Opa cluster
type Client struct {
	client   *http.Client
	baseURL  string
	common   service
	Policies *PolicyService

	logger io.Writer
	debug  bool

	custom.Registry
}

// Status respresents current status of a Opa node.
type Status struct {
	Database struct {
		Reachable bool `json:"reachable"`
	} `json:"database"`
	Server struct {
		ConnectionsAccepted int `json:"connections_accepted"`
		ConnectionsActive   int `json:"connections_active"`
		ConnectionsHandled  int `json:"connections_handled"`
		ConnectionsReading  int `json:"connections_reading"`
		ConnectionsWaiting  int `json:"connections_waiting"`
		ConnectionsWriting  int `json:"connections_writing"`
		TotalRequests       int `json:"total_requests"`
	} `json:"server"`
}

// NewClient returns a Client which talks to Admin API of Opa
func NewClient(baseURL *string, client *http.Client) (*Client, error) {
	if client == nil {
		client = http.DefaultClient
	}
	Opa := new(Client)
	Opa.client = client
	var rootURL string
	if baseURL != nil {
		rootURL = *baseURL
	} else {
		rootURL = defaultBaseURL
	}
	url, err := url.ParseRequestURI(rootURL)
	if err != nil {
		return nil, errors.Wrap(err, "parsing URL")
	}
	Opa.baseURL = url.String()

	Opa.common.client = Opa
	Opa.Policies = (*PolicyService)(&Opa.common)
	Opa.Registry = custom.NewDefaultRegistry()

	Opa.logger = os.Stderr
	return Opa, nil
}

// Do executes a HTTP request and returns a response
func (c *Client) Do(ctx context.Context, req *http.Request,
	v interface{}) (*Response, error) {
	var err error
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}
	if ctx == nil {
		ctx = defaultCtx
	}
	req = req.WithContext(ctx)

	// log the request
	c.logRequest(req)

	//Make the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "making HTTP request")
	}

	// log the response
	c.logResponse(resp)

	response := newResponse(resp)

	///check for API errors
	if err = hasError(resp); err != nil {
		return response, err
	}
	// Call Close on exit
	defer func() {
		e := resp.Body.Close()
		if e != nil {
			err = e
		}
	}()

	// response
	if v != nil {
		if writer, ok := v.(io.Writer); ok {
			_, err = io.Copy(writer, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return response, err
}

// SetDebugMode enables or disables logging of
// the request to the logger set by SetLogger().
// By default, debug logging is disabled.
func (c *Client) SetDebugMode(enableDebug bool) {
	c.debug = enableDebug
}

func (c *Client) logRequest(r *http.Request) error {
	if !c.debug {
		return nil
	}
	dump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		return err
	}
	_, err = c.logger.Write(append(dump, '\n'))
	return err
}

func (c *Client) logResponse(r *http.Response) error {
	if !c.debug {
		return nil
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		return err
	}
	_, err = c.logger.Write(append(dump, '\n'))
	return err
}

// SetLogger sets the debug logger, defaults to os.StdErr
func (c *Client) SetLogger(w io.Writer) {
	if w == nil {
		return
	}
	c.logger = w
}

// Status returns the status of a Opa node
func (c *Client) Status(ctx context.Context) (*Status, error) {

	req, err := c.NewJsonRequest("GET", "/status", nil, nil)
	if err != nil {
		return nil, err
	}

	var s Status
	_, err = c.Do(ctx, req, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// Root returns the response of GET request on root of
// Admin API (GET /).
func (c *Client) Root(ctx context.Context) (map[string]interface{}, error) {
	req, err := c.NewJsonRequest("GET", "/", nil, nil)
	if err != nil {
		return nil, err
	}

	var root map[string]interface{}
	_, err = c.Do(ctx, req, &root)
	if err != nil {
		return nil, err
	}
	return root, nil
}

// Root returns the response of GET request on root of
// Admin API (GET /).
func (c *Client) Health(ctx context.Context) (map[string]interface{}, error) {
	req, err := c.NewJsonRequest("GET", "/health", nil, nil)
	if err != nil {
		return nil, err
	}

	var root map[string]interface{}
	_, err = c.Do(ctx, req, &root)
	if err != nil {
		return nil, err
	}
	return root, nil
}