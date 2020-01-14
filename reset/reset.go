package reset

import (
	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
)

// Reset deletes all entities in Opa.
func Reset(state *utils.OpaRawState, client *opa.Client) error {
	if state == nil {
		return errors.New("state cannot be empty")
	}
	// TODO parallelize these operations

	// Delete policies
	for _, r := range state.Policies {
		err := client.Policies.Delete(nil, r.ID)
		if err != nil {
			return err
		}
	}

	// TODO handle custom entities
	return nil
}
