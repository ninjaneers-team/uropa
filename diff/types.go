package diff

import (
	"github.com/ninjaneers-team/uropa/crud"
)

// Event represents an event to perform
// an imperative operation
// that gets Opa closer to the target state.
type Event struct {
	Op     crud.Op
	Kind   crud.Kind
	Obj    interface{}
	OldObj interface{}
}
