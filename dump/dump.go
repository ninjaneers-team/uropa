package dump

import (
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
)

// Config can be used to skip exporting certain entities
type Config struct {
	// If true, consumers and any plugins associated with it
	// are not exported.
	SkipConsumers bool

	// SelectorTags can be used to export entities tagged with only specific
	// tags.
	SelectorTags []string
}

func newOpt(tags []string) *opa.ListOpt {
	opt := new(opa.ListOpt)
	opt.Size = 1000
	opt.Tags = opa.StringSlice(tags...)
	opt.MatchAllTags = true
	return opt
}

// Get queries all the entities using client and returns
// all the entities in OpaRawState.
func Get(client *opa.Client, config Config) (*utils.OpaRawState, error) {

	// TODO make these requests concurrent

	var state utils.OpaRawState
	policies, err := GetAllPolicies(client, config.SelectorTags)
	if err != nil {
		return nil, errors.Wrap(err, "policies")
	}

	state.Policies = policies

	return &state, nil
}

// GetAllPolicies queries Opa for all the policies using client.
func GetAllPolicies(client *opa.Client, tags []string) ([]*opa.Policy, error) {
	var policies []*opa.Policy
	opt := newOpt(tags)

	for {
		s, nextopt, err := client.Policies.List(nil, opt)
		if err != nil {
			return nil, err
		}
		policies = append(policies, s...)
		if nextopt == nil {
			break
		}
		opt = nextopt
	}
	return policies, nil
}
