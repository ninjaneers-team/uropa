package state

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/pkg/errors"
)

type collection struct {
	db *memdb.MemDB
}

// OpaState is an in-memory database representation
// of Opa's configuration.
type OpaState struct {
	common   collection
	Policies *PoliciesCollection
}

// NewOpaState creates a new in-memory OpaState.
func NewOpaState() (*OpaState, error) {

	var schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			policyTableName: policyTableSchema,
		},
	}

	memDB, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, errors.Wrap(err, "creating new ServiceCollection")
	}
	var state OpaState
	state.common = collection{
		db: memDB,
	}

	state.Policies = (*PoliciesCollection)(&state.common)

	return &state, nil
}
