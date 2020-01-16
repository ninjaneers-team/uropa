package opa

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// PolicyService handles Policies in Opa.
type PolicyService service

// Create creates a Policy in Opa.
// If an ID is specified, it will be used to
// create a policy in Opa, otherwise an ID
// is auto-generated.
func (s *PolicyService) Create(ctx context.Context,
	policy *Policy) (*Policy, error) {

	queryPath := "/v1/policies"
	queryPath = queryPath + "/" + *policy.ID
	method := "PUT"

	handler, err := s.client.NewFile(*policy.Raw)

	if err != nil {
		return nil, err
	}

	f, err := os.Open(handler.Name())
	if err != nil {
		panic(err)
	}

	req, err := s.client.NewFormEncodedRequest(method, queryPath, nil, f)

	defer f.Close()
	defer os.Remove(f.Name())
	defer handler.Close()

	var createdPolicy Policy
	_, err = s.client.Do(ctx, req, &createdPolicy)
	if err != nil {
		return nil, err
	}
	return &createdPolicy, nil
}

// Get fetches a Policy in Opa.
func (s *PolicyService) Get(ctx context.Context,
	usernameOrID *string) (*Policy, error) {

	if isEmptyString(usernameOrID) {
		return nil, errors.New("usernameOrID cannot be nil for Get operation")
	}

	endpoint := fmt.Sprintf("/v1/policies/%v", *usernameOrID)
	req, err := s.client.NewJsonRequest("GET", endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var policy Policy
	_, err = s.client.Do(ctx, req, &policy)
	if err != nil {
		return nil, err
	}
	return &policy, nil
}

// GetByCustomID fetches a Policy in Opa.
func (s *PolicyService) GetByCustomID(ctx context.Context,
	customID *string) (*Policy, error) {

	if isEmptyString(customID) {
		return nil, errors.New("customID cannot be nil for Get operation")
	}

	type QS struct {
		CustomID string `url:"custom_id,omitempty"`
	}

	req, err := s.client.NewJsonRequest("GET", "/v1/policies",
		&QS{CustomID: *customID}, nil)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Data []Policy
	}
	var resp Response
	_, err = s.client.Do(ctx, req, &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, err404{}
	}

	return &resp.Data[0], nil
}

// Update updates a Policy in Opa
func (s *PolicyService) Update(ctx context.Context,
	policy *Policy) (*Policy, error) {
	return s.Create(ctx, policy)
}

// Delete deletes a Policy in Opa
func (s *PolicyService) Delete(ctx context.Context,
	usernameOrID *string) error {

	if isEmptyString(usernameOrID) {
		return errors.New("usernameOrID cannot be nil for Delete operation")
	}

	endpoint := fmt.Sprintf("/v1/policies/%v", *usernameOrID)
	req, err := s.client.NewJsonRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(ctx, req, nil)
	return err
}

// List fetches a list of Policies in Opa.
// opt can be used to control pagination.
func (s *PolicyService) List(ctx context.Context,
	opt *ListOpt) ([]*Policy, *ListOpt, error) {
	data, next, err := s.client.list(ctx, "/v1/policies", opt)
	if err != nil {
		return nil, nil, err
	}
	var policies []*Policy

	for _, object := range data {
		b, err := object.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		var policy Policy
		err = json.Unmarshal(b, &policy)
		if err != nil {
			return nil, nil, err
		}

		policies = append(policies, &policy)
	}

	return policies, next, nil
}

// ListAll fetches all Policies in Opa.
// This method can take a while if there
// a lot of Policies present.
func (s *PolicyService) ListAll(ctx context.Context) ([]*Policy, error) {
	var policies, data []*Policy
	var err error
	opt := &ListOpt{Pretty: pretty}

	for opt != nil {
		data, opt, err = s.List(ctx, opt)
		if err != nil {
			return nil, err
		}
		policies = append(policies, data...)
	}
	return policies, nil
}
