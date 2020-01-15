package diff

import (
	"github.com/ninjaneers-team/uropa/crud"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/pkg/errors"
)

func (sc *Syncer) deletePolicies() error {
	currentPolicies, err := sc.currentState.Policies.GetAll()
	if err != nil {
		return errors.Wrap(err, "error fetching policies from state")
	}

	for _, policy := range currentPolicies {
		n, err := sc.deletePolicy(policy)
		if err != nil {
			return err
		}
		if n != nil {
			err = sc.queueEvent(*n)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (sc *Syncer) deletePolicy(policy *state.Policy) (*Event, error) {
	_, err := sc.targetState.Policies.Get(*policy.ID)
	if err == state.ErrNotFound {
		return &Event{
			Op:   crud.Delete,
			Kind: "policy",
			Obj:  policy,
		}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "looking up policy '%v'",
			policy.Identifier())
	}
	return nil, nil
}

func (sc *Syncer) createUpdatePolicies() error {
	targetPolicies, err := sc.targetState.Policies.GetAll()
	if err != nil {
		return errors.Wrap(err, "error fetching policies from state")
	}

	for _, policy := range targetPolicies {
		n, err := sc.createUpdatePolicy(policy)
		if err != nil {
			return err
		}
		if n != nil {
			err = sc.queueEvent(*n)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (sc *Syncer) createUpdatePolicy(policy *state.Policy) (*Event, error) {
	policyCopy := &state.Policy{Policy: *policy.DeepCopy()}
	currentPolicy, err := sc.currentState.Policies.Get(*policy.ID)

	if err == state.ErrNotFound {
		return &Event{
			Op:   crud.Create,
			Kind: "policy",
			Obj:  policyCopy,
		}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error looking up policy %v",
			*policy.ID)
	}

	// found, check if update needed
	if !currentPolicy.EqualWithOpts(policyCopy, false, true) {
		return &Event{
			Op:     crud.Update,
			Kind:   "policy",
			Obj:    policyCopy,
			OldObj: currentPolicy,
		}, nil
	}
	return nil, nil
}
