package file

import (
	"encoding/hex"
	"math/rand"
	"os"
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/stretchr/testify/assert"
)

func emptyState() *state.OpaState {
	s, _ := state.NewOpaState()
	return s
}

func existingPolicyState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Policies.Add(state.Policy{
		Policy: opa.Policy{
			ID:  opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Raw: opa.String("foo"),
		},
	})
	return s
}

var deterministicUUID = func() *string {
	version := byte(4)
	uuid := make([]byte, 16)
	rand.Read(uuid)

	// Set version
	uuid[6] = (uuid[6] & 0x0f) | (version << 4)

	// Set variant
	uuid[8] = (uuid[8] & 0xbf) | 0x80

	buf := make([]byte, 36)
	var dash byte = '-'
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], uuid[10:])
	s := string(buf)
	return &s
}

func TestMain(m *testing.M) {
	uuid = deterministicUUID
	os.Exit(m.Run())
}

func Test_stateBuilder_policies(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		targetContent *Content
		currentState  *state.OpaState
	}
	tests := []struct {
		name   string
		fields fields
		want   *utils.OpaRawState
	}{
		{
			name: "matches ID of an existing policy",
			fields: fields{
				targetContent: &Content{
					Policies: []FPolicy{
						{
							Policy: opa.Policy{
								ID: opa.String("foo"),
							},
						},
					},
				},
				currentState: existingPolicyState(),
			},
			want: &utils.OpaRawState{
				Policies: []*opa.Policy{
					{
						Raw: opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						ID:  opa.String("foo"),
					},
				},
			},
		},
		{
			name: "process a non-existent policy",
			fields: fields{
				targetContent: &Content{
					Policies: []FPolicy{
						{
							Policy: opa.Policy{
								ID: opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
							},
						},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				Policies: []*opa.Policy{
					{
						ID:  opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Raw: opa.String("foo"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &stateBuilder{
				targetContent: tt.fields.targetContent,
				currentState:  tt.fields.currentState,
			}
			d, _ := utils.GetOpaDefaulter()
			b.defaulter = d
			b.build()
			assert.Equal(tt.want, b.rawState)
		})
	}
}
