package state

import (
	"github.com/ninjaneers-team/uropa/opa"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func policysCollection() *PoliciesCollection {
	return state().Policies
}

func TestPolicysCollection_Add(t *testing.T) {
	type args struct {
		policy Policy
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "errors when ID is nil",
			args: args{
				policy: Policy{
					Policy: opa.Policy{
						Raw: opa.String("example.com"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "errors on re-insert by ID",
			args: args{
				policy: Policy{
					Policy: opa.Policy{
						ID:  opa.String("id3"),
						Raw: opa.String("example.com"),
					},
				},
			},
			wantErr: true,
		},
	}
	k := policysCollection()
	svc1 := Policy{
		Policy: opa.Policy{
			ID:  opa.String("id3"),
			Raw: opa.String("example.com"),
		},
	}
	k.Add(svc1)
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := k.Add(tt.args.policy); (err != nil) != tt.wantErr {
				t.Errorf("PoliciesCollection.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPolicysCollection_Get(t *testing.T) {
	type args struct {
		nameOrID string
	}
	svc1 := Policy{
		Policy: opa.Policy{
			ID:  opa.String("foo-id"),
			Raw: opa.String("example.com"),
		},
	}
	svc2 := Policy{
		Policy: opa.Policy{
			ID:  opa.String("bar-id"),
			Raw: opa.String("example.com"),
		},
	}
	tests := []struct {
		name    string
		args    args
		want    *Policy
		wantErr bool
	}{

		{
			name: "gets a policy by ID",
			args: args{
				nameOrID: "foo-id",
			},
			want:    &svc1,
			wantErr: false,
		},
		{
			name: "gets a policy by Name",
			args: args{
				nameOrID: "bar-name",
			},
			want:    &svc2,
			wantErr: false,
		},
		{
			name: "returns an ErrNotFound when no policy found",
			args: args{
				nameOrID: "baz-id",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error when ID is empty",
			args: args{
				nameOrID: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	k := policysCollection()
	k.Add(svc1)
	k.Add(svc2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := k.Get(tt.args.nameOrID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PoliciesCollection.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PoliciesCollection.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolicysCollection_Update(t *testing.T) {
	pol1 := Policy{
		Policy: opa.Policy{
			ID:  opa.String("foo-id"),
			Raw: opa.String("example.com"),
		},
	}
	pol2 := Policy{
		Policy: opa.Policy{
			ID:  opa.String("bar-id"),
			Raw: opa.String("example.com"),
		},
	}
	svc3 := Policy{
		Policy: opa.Policy{
			ID:  opa.String("foo-id"),
			Raw: opa.String("2.example.com"),
		},
	}
	type args struct {
		policy Policy
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		updatedPolicy *Policy
	}{
		{
			name: "update errors if policy.ID is nil",
			args: args{
				policy: Policy{
					Policy: opa.Policy{
						Raw: opa.String("name"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "update errors if policy does not exist",
			args: args{
				policy: Policy{
					Policy: opa.Policy{
						ID: opa.String("does-not-exist"),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "update succeeds when ID is supplied",
			args: args{
				policy: svc3,
			},
			wantErr:       false,
			updatedPolicy: &svc3,
		},
	}
	k := policysCollection()
	k.Add(pol1)
	k.Add(pol2)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			if err := k.Update(tt.args.policy); (err != nil) != tt.wantErr {
				t.Errorf("PoliciesCollection.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				got, _ := k.Get(*tt.updatedPolicy.ID)

				if !reflect.DeepEqual(got, tt.updatedPolicy) {
					t.Errorf("update policy, got = %#v, want %#v", got, tt.updatedPolicy)
				}
			}
		})
	}
}

func TestPolicyDelete(t *testing.T) {
	assert := assert.New(t)
	collection := policysCollection()

	var policy Policy
	policy.ID = opa.String("first")
	policy.Raw = opa.String("example.com")
	err := collection.Add(policy)
	assert.Nil(err)

	err = collection.Delete("does-not-exist")
	assert.NotNil(err)
	err = collection.Delete("first")
	assert.Nil(err)

	err = collection.Delete("first")
	assert.NotNil(err)

	err = collection.Delete("")
	assert.NotNil(err)
}

func TestPolicyGetAll(t *testing.T) {
	assert := assert.New(t)
	collection := policysCollection()

	policys := []Policy{
		{
			Policy: opa.Policy{
				ID:  opa.String("first"),
				Raw: opa.String("example.com"),
			},
		},
		{
			Policy: opa.Policy{
				ID:  opa.String("second"),
				Raw: opa.String("example.com"),
			},
		},
	}
	for _, s := range policys {
		assert.Nil(collection.Add(s))
	}

	allPolicys, err := collection.GetAll()

	assert.Nil(err)
	assert.Equal(len(policys), len(allPolicys))
}
