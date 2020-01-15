package state

import (
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
)

// Get builds a KongState from a raw representation of Opa.
func Get(raw *utils.OpaRawState) (*OpaState, error) {
	opaState, err := NewOpaState()
	if err != nil {
		return nil, errors.Wrap(err, "creating new in-memory state of Opa")
	}

	for _, s := range raw.Policies {
		err := opaState.Policies.Add(Policy{Policy: *s})
		if err != nil {
			return nil, errors.Wrap(err, "inserting service into state")
		}
	}

	return opaState, nil
}
