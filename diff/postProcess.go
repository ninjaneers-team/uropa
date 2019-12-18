package diff

import (
	"github.com/ninjaneers-team/uropa/crud"
	"github.com/ninjaneers-team/uropa/state"
)

type policyPostAction struct {
	currentState *state.OpaState
}

func (crud *policyPostAction) Create(args ...crud.Arg) (crud.Arg, error) {
	return nil, crud.currentState.Policies.Add(*args[0].(*state.Policy))
}

func (crud *policyPostAction) Delete(args ...crud.Arg) (crud.Arg, error) {
	return nil, crud.currentState.Policies.Delete(*((args[0].(*state.Policy)).ID))
}

func (crud *policyPostAction) Update(args ...crud.Arg) (crud.Arg, error) {
	return nil, crud.currentState.Policies.Update(*args[0].(*state.Policy))
}
