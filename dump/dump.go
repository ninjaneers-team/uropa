package dump

import (
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
)

// Config can be used to skip exporting certain entities
type Config struct {
	Pretty bool
}

func newOpt() *opa.ListOpt {
	opt := new(opa.ListOpt)
	opt.Pretty = false
	return opt
}

// Get queries all the entities using client and returns
// all the entities in OpaRawState.
func Get(client *opa.Client, config Config) (*utils.OpaRawState, error) {
	var state utils.OpaRawState
	policies, err := GetAllPolicies(client)
	if err != nil {
		return nil, errors.Wrap(err, "policies")
	}

	state.Policies = policies

	return &state, nil
}

// GetAllPolicies queries Opa for all the policies using client.
func GetAllPolicies(client *opa.Client) ([]*opa.Policy, error) {
	var policies []*opa.Policy
	opt := newOpt()

	for {
		s, nextopt, err := client.Policies.List(nil, opt)
		if err != nil {
			return nil, err
		}
		policies = append(policies, s...)
		if nextopt == nil {
			break
		}
	}
	return policies, nil
}
