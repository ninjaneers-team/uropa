package opa

import (
	"context"
	"encoding/json"
)

// ListOpt aids in paginating through list endpoints
type ListOpt struct {
	Pretty bool
}

// qs is used to construct query string for list endpoints
type qs struct {
	Pretty bool `url:"pretty,omitempty"`
}

// list fetches a list of an entity in Opa.
// opt can be used to control pagination.
func (c *Client) list(ctx context.Context,
	endpoint string, opt *ListOpt) ([]json.RawMessage, *ListOpt, error) {

	q := constructQueryString(opt)
	req, err := c.NewJSONRequest("GET", endpoint, &q, nil)
	if err != nil {
		return nil, nil, err
	}
	var list struct {
		Data []json.RawMessage `json:"result"`
		Next *string           `json:"offset"`
	}

	_, err = c.Do(ctx, req, &list)
	if err != nil {
		return nil, nil, err
	}

	// convinient for end user to use this opt till it's nil
	var next *ListOpt
	if list.Next != nil {
	}

	return list.Data, next, nil
}

func constructQueryString(opt *ListOpt) qs {
	var q qs
	if opt == nil {
		return q
	}
	q.Pretty = opt.Pretty

	return q
}
