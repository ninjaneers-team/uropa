package state

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/ninjaneers-team/uropa/utils"
)

const (
	policyTableName = "policy"
)

var policyTableSchema = &memdb.TableSchema{
	Name: policyTableName,
	Indexes: map[string]*memdb.IndexSchema{
		"id": {
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "ID"},
		},
		all: allIndex,
	},
}

// PoliciesCollection stores and indexes Opa Policies.
type PoliciesCollection collection

// Add adds a policy to the collection.
// policy.ID should not be nil else an error is thrown.
func (k *PoliciesCollection) Add(policy Policy) error {
	if utils.Empty(policy.ID) {
		return errIDRequired
	}
	txn := k.db.Txn(true)
	defer txn.Abort()

	var searchBy []string
	searchBy = append(searchBy, *policy.ID)
	if !utils.Empty(policy.ID) {
		searchBy = append(searchBy, *policy.ID)
	}
	_, err := getPolicy(txn, searchBy...)
	if err == nil {
		return ErrAlreadyExists
	} else if err != ErrNotFound {
		return err
	}

	err = txn.Insert(policyTableName, &policy)
	if err != nil {
		return err
	}
	txn.Commit()
	return nil
}

func getPolicy(txn *memdb.Txn, IDs ...string) (*Policy, error) {
	for _, id := range IDs {
		res, err := multiIndexLookupUsingTxn(txn, policyTableName,
			[]string{"id"}, id)
		if err == ErrNotFound {
			continue
		}
		if err != nil {
			return nil, err
		}
		policy, ok := res.(*Policy)
		if !ok {
			panic(unexpectedType)
		}
		return &Policy{Policy: *policy.DeepCopy()}, nil
	}
	return nil, ErrNotFound

}

// Get gets a policy by name or ID.
func (k *PoliciesCollection) Get(id string) (*Policy, error) {
	if id == "" {
		return nil, errIDRequired
	}

	txn := k.db.Txn(false)
	defer txn.Abort()
	return getPolicy(txn, id)
}

// Update udpates an existing policy.
// It returns an error if the policy is not already present.
func (k *PoliciesCollection) Update(policy Policy) error {
	// TODO abstract this check in the go-memdb library itself
	if utils.Empty(policy.ID) {
		return errIDRequired
	}

	txn := k.db.Txn(true)
	defer txn.Abort()

	err := deletePolicy(txn, *policy.ID)
	if err != nil {
		return err
	}

	err = txn.Insert(policyTableName, &policy)
	if err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func deletePolicy(txn *memdb.Txn, nameOrID string) error {
	policy, err := getPolicy(txn, nameOrID)
	if err != nil {
		return err
	}

	err = txn.Delete(policyTableName, policy)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a policy by name or ID.
func (k *PoliciesCollection) Delete(id string) error {
	if id == "" {
		return errIDRequired
	}

	txn := k.db.Txn(true)
	defer txn.Abort()

	err := deletePolicy(txn, id)
	if err != nil {
		return err
	}

	txn.Commit()
	return nil
}

// GetAll returns all the policies.
func (k *PoliciesCollection) GetAll() ([]*Policy, error) {
	txn := k.db.Txn(false)
	defer txn.Abort()

	iter, err := txn.Get(policyTableName, all, true)
	if err != nil {
		return nil, err
	}

	var res []*Policy
	for el := iter.Next(); el != nil; el = iter.Next() {
		s, ok := el.(*Policy)
		if !ok {
			panic(unexpectedType)
		}
		res = append(res, &Policy{Policy: *s.DeepCopy()})
	}
	txn.Commit()
	return res, nil
}
