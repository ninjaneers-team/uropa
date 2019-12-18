package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

// OpaRawState contains all of Opa Data
type OpaRawState struct {
	Policies []*opa.Policy
}

// ErrArray holds an array of errors.
type ErrArray struct {
	Errors []error
}

// Error returns a pretty string of errors present.
func (e ErrArray) Error() string {
	if len(e.Errors) == 0 {
		return "nil"
	}
	var res string

	res = strconv.Itoa(len(e.Errors)) + " errors occurred:\n"
	for _, err := range e.Errors {
		res += fmt.Sprintf("\t%v\n", err)
	}
	return res
}

// OpaClientConfig holds config details to use to talk to a Opa server.
type OpaClientConfig struct {
	Address   string
	Workspace string

	Headers []string

	TLSSkipVerify bool
	TLSServerName string

	TLSCACert string

	Debug bool
}

// HeaderRoundTripper injects Headers into requests
// made via RT.
type HeaderRoundTripper struct {
	headers []string
	rt      http.RoundTripper
}

// RoundTrip satisfies the RoundTripper interface.
func (t *HeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response,
	error) {
	newRequest := new(http.Request)
	*newRequest = *req
	newRequest.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		newRequest.Header[k] = append([]string(nil), s...)
	}
	for _, s := range t.headers {
		split := strings.SplitN(s, ":", 2)
		if len(split) >= 2 {
			newRequest.Header[split[0]] = append([]string(nil), split[1])
		}
	}
	return t.rt.RoundTrip(newRequest)
}

// GetOpaClient returns a Opa opa
func GetOpaClient(opt OpaClientConfig) (*opa.Client, error) {

	var tlsConfig tls.Config
	if opt.TLSSkipVerify {
		tlsConfig.InsecureSkipVerify = true
	}
	if opt.TLSServerName != "" {
		tlsConfig.ServerName = opt.TLSServerName
	}

	if opt.TLSCACert != "" {
		certPool := x509.NewCertPool()
		ok := certPool.AppendCertsFromPEM([]byte(opt.TLSCACert))
		if !ok {
			return nil, errors.New("failed to load TLSCACert")
		}
		tlsConfig.RootCAs = certPool
	}

	c := &http.Client{}
	defaultTransport := http.DefaultTransport.(*http.Transport)
	defaultTransport.TLSClientConfig = &tlsConfig
	c.Transport = defaultTransport
	if len(opt.Headers) > 0 {
		c.Transport = &HeaderRoundTripper{
			headers: opt.Headers,
			rt:      defaultTransport,
		}
	}

	url, err := url.Parse(opt.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse Opa address")
	}
	if opt.Workspace != "" {
		url.Path = path.Join(url.Path, opt.Workspace)
	}

	OpaClient, err := opa.NewClient(opa.String(url.String()), c)
	if err != nil {
		return nil, errors.Wrap(err, "creating opa for Opa's Admin API")
	}
	if opt.Debug {
		OpaClient.SetDebugMode(true)
		OpaClient.SetLogger(os.Stderr)
	}
	return OpaClient, nil
}
