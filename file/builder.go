package file

import (
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/ninjaneers-team/uropa/utils"
)

type stateBuilder struct {
	targetContent *Content
	rawState      *utils.OpaRawState
	currentState  *state.OpaState
	defaulter     *utils.Defaulter

	selectTags   []string
	intermediate *state.OpaState
	err          error
}

// uuid generates a UUID string and returns a pointer to it.
// It is a variable for testing purpose, to override and supply
// a deterministic UUID generator.
var uuid = func() *string {
	return opa.String(utils.UUID())
}

func (b *stateBuilder) build() (*utils.OpaRawState, error) {
	// setup
	var err error
	b.rawState = &utils.OpaRawState{}

	if b.targetContent.Info != nil {
		b.selectTags = b.targetContent.Info.SelectorTags
	}
	b.intermediate, err = state.NewOpaState()
	if err != nil {
		return nil, err
	}

	// build
	b.policies()

	// result
	if b.err != nil {
		return nil, b.err
	}
	return b.rawState, nil
}

func (b *stateBuilder) policies() {
	if b.err != nil {
		return
	}

	for _, r := range b.targetContent.Policies {
		r := r
		if err := b.ingestPolicy(r); err != nil {
			b.err = err
			return
		}
	}
}

func (b *stateBuilder) ingestPolicy(r FPolicy) error {
	if utils.Empty(r.ID) {
		policy, err := b.currentState.Policies.Get(*r.ID)
		if err == state.ErrNotFound {
			r.ID = uuid()
		} else if err != nil {
			return err
		} else {
			r.ID = opa.String(*policy.ID)
		}
	}

	b.rawState.Policies = append(b.rawState.Policies, &r.Policy)
	err := b.intermediate.Policies.Add(state.Policy{Policy: r.Policy})
	if err != nil {
		return err
	}

	return nil
}
