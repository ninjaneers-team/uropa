package solver

import (
	"github.com/ninjaneers-team/uropa/crud"
	"github.com/ninjaneers-team/uropa/diff"
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/state"
)

// policyCRUD implements crud.Actions interface.
type policyCRUD struct {
	client *opa.Client
}

func policyFromStuct(arg diff.Event) *state.Policy {
	policy, ok := arg.Obj.(*state.Policy)
	if !ok {
		panic("unexpected type, expected *state.policy")
	}
	return policy
}

// Create a Policy in Opa.
// The arg should be of type diff.Event, containing the policy to be created,
// else the function will panic.
// It returns a the created *state.Policy.
func (s *policyCRUD) Create(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	policy := policyFromStuct(event)
	_, err := s.client.Policies.Create(nil, &policy.Policy)
	if err != nil {
		return nil, err
	}
	return &state.Policy{Policy: policy.Policy}, nil
}

// Delete deletes a Policy in Opa.
// The arg should be of type diff.Event, containing the policy to be deleted,
// else the function will panic.
// It returns a the deleted *state.Policy.
func (s *policyCRUD) Delete(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	policy := policyFromStuct(event)
	err := s.client.Policies.Delete(nil, policy.ID)
	if err != nil {
		return nil, err
	}
	return policy, nil
}

// Update updates a Policy in Opa.
// The arg should be of type diff.Event, containing the policy to be updated,
// else the function will panic.
// It returns a the updated *state.Policy.
func (s *policyCRUD) Update(arg ...crud.Arg) (crud.Arg, error) {
	event := eventFromArg(arg[0])
	policy := policyFromStuct(event)

	updatedPolicy, err := s.client.Policies.Create(nil, &policy.Policy)
	if err != nil {
		return nil, err
	}
	return &state.Policy{Policy: *updatedPolicy}, nil
}
