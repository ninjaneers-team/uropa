package file

import (
	"encoding/hex"
	"math/rand"
	"os"
	"reflect"
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

func existingRouteState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Routes.Add(state.Route{
		Route: Opa.Route{
			ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Name: Opa.String("foo"),
		},
	})
	return s
}

func existingServiceState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Services.Add(state.Service{
		Service: Opa.Service{
			ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Name: Opa.String("foo"),
		},
	})
	return s
}

func existingConsumerCredState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Consumers.Add(state.Consumer{
		Consumer: Opa.Consumer{
			ID:       Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Username: Opa.String("foo"),
		},
	})
	s.KeyAuths.Add(state.KeyAuth{
		KeyAuth: Opa.KeyAuth{
			ID:  Opa.String("5f1ef1ea-a2a5-4a1b-adbb-b0d3434013e5"),
			Key: Opa.String("foo-apikey"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			},
		},
	})
	s.BasicAuths.Add(state.BasicAuth{
		BasicAuth: Opa.BasicAuth{
			ID:       Opa.String("92f4c849-960b-43af-aad3-f307051408d3"),
			Username: Opa.String("basic-username"),
			Password: Opa.String("basic-password"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			},
		},
	})
	s.JWTAuths.Add(state.JWTAuth{
		JWTAuth: Opa.JWTAuth{
			ID:     Opa.String("917b9402-1be0-49d2-b482-ca4dccc2054e"),
			Key:    Opa.String("jwt-key"),
			Secret: Opa.String("jwt-secret"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			},
		},
	})
	s.HMACAuths.Add(state.HMACAuth{
		HMACAuth: Opa.HMACAuth{
			ID:       Opa.String("e5d81b73-bf9e-42b0-9d68-30a1d791b9c9"),
			Username: Opa.String("hmac-username"),
			Secret:   Opa.String("hmac-secret"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			},
		},
	})
	s.ACLGroups.Add(state.ACLGroup{
		ACLGroup: Opa.ACLGroup{
			ID:    Opa.String("b7c9352a-775a-4ba5-9869-98e926a3e6cb"),
			Group: Opa.String("foo-group"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			},
		},
	})
	s.Oauth2Creds.Add(state.Oauth2Credential{
		Oauth2Credential: Opa.Oauth2Credential{
			ID:       Opa.String("4eef5285-3d6a-4f6b-b659-8957a940e2ca"),
			ClientID: Opa.String("oauth2-clientid"),
			Name:     Opa.String("oauth2-name"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			},
		},
	})
	return s
}

func existingUpstreamState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Upstreams.Add(state.Upstream{
		Upstream: Opa.Upstream{
			ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Name: Opa.String("foo"),
		},
	})
	return s
}

func existingCertificateState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Certificates.Add(state.Certificate{
		Certificate: Opa.Certificate{
			ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Cert: Opa.String("foo"),
			Key:  Opa.String("bar"),
		},
	})
	return s
}

func existingCACertificateState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.CACertificates.Add(state.CACertificate{
		CACertificate: Opa.CACertificate{
			ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Cert: Opa.String("foo"),
		},
	})
	return s
}

func existingPluginState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Plugins.Add(state.Plugin{
		Plugin: Opa.Plugin{
			ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
			Name: Opa.String("foo"),
		},
	})
	s.Plugins.Add(state.Plugin{
		Plugin: Opa.Plugin{
			ID:   Opa.String("f7e64af5-e438-4a9b-8ff8-ec6f5f06dccb"),
			Name: Opa.String("bar"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
			},
		},
	})
	s.Plugins.Add(state.Plugin{
		Plugin: Opa.Plugin{
			ID:   Opa.String("53ce0a9c-d518-40ee-b8ab-1ee83a20d382"),
			Name: Opa.String("foo"),
			Consumer: &Opa.Consumer{
				ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
			},
			Route: &Opa.Route{
				ID: Opa.String("700bc504-b2b1-4abd-bd38-cec92779659e"),
			},
		},
	})
	return s
}

func existingTargetsState() *state.OpaState {
	s, _ := state.NewOpaState()
	s.Targets.Add(state.Target{
		Target: Opa.Target{
			ID:     Opa.String("f7e64af5-e438-4a9b-8ff8-ec6f5f06dccb"),
			Target: Opa.String("bar"),
			Upstream: &Opa.Upstream{
				ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
			},
		},
	})
	s.Targets.Add(state.Target{
		Target: Opa.Target{
			ID:     Opa.String("53ce0a9c-d518-40ee-b8ab-1ee83a20d382"),
			Target: Opa.String("foo"),
			Upstream: &Opa.Upstream{
				ID: Opa.String("700bc504-b2b1-4abd-bd38-cec92779659e"),
			},
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

func Test_stateBuilder_services(t *testing.T) {
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
					Info: &Info{
						SelectorTags: []string{"tag1"},
					},
					Services: []FPolicy{
						{
							Service: Opa.Service{
								Name: Opa.String("foo"),
							},
						},
					},
				},
				currentState: existingServiceState(),
			},
			want: &utils.OpaRawState{
				Services: []*Opa.Service{
					{
						ID:             Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Name:           Opa.String("foo"),
						Port:           Opa.Int(80),
						Retries:        Opa.Int(5),
						Protocol:       Opa.String("http"),
						ConnectTimeout: Opa.Int(60000),
						WriteTimeout:   Opa.Int(60000),
						ReadTimeout:    Opa.Int(60000),
						Tags:           Opa.StringSlice("tag1"),
					},
				},
			},
		},
		{
			name: "process a non-existent policy",
			fields: fields{
				targetContent: &Content{
					Services: []FPolicy{
						{
							Service: Opa.Service{
								Name: Opa.String("foo"),
							},
						},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				Services: []*Opa.Service{
					{
						ID:             Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Name:           Opa.String("foo"),
						Port:           Opa.Int(80),
						Retries:        Opa.Int(5),
						Protocol:       Opa.String("http"),
						ConnectTimeout: Opa.Int(60000),
						WriteTimeout:   Opa.Int(60000),
						ReadTimeout:    Opa.Int(60000),
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

func Test_stateBuilder_ingestRoute(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		currentState *state.OpaState
	}
	type args struct {
		route FRoute
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantState *utils.OpaRawState
	}{
		{
			name: "generates ID for a non-existing route",
			fields: fields{
				currentState: emptyState(),
			},
			args: args{
				route: FRoute{
					Route: Opa.Route{
						Name: Opa.String("foo"),
					},
				},
			},
			wantErr: false,
			wantState: &utils.OpaRawState{
				Routes: []*Opa.Route{
					{
						ID:            Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Name:          Opa.String("foo"),
						PreserveHost:  Opa.Bool(false),
						RegexPriority: Opa.Int(0),
						StripPath:     Opa.Bool(false),
						Protocols:     Opa.StringSlice("http", "https"),
					},
				},
			},
		},
		{
			name: "matches up IDs of routes correctly",
			fields: fields{
				currentState: existingRouteState(),
			},
			args: args{
				route: FRoute{
					Route: Opa.Route{
						Name: Opa.String("foo"),
					},
				},
			},
			wantErr: false,
			wantState: &utils.OpaRawState{
				Routes: []*Opa.Route{
					{
						ID:            Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Name:          Opa.String("foo"),
						PreserveHost:  Opa.Bool(false),
						RegexPriority: Opa.Int(0),
						StripPath:     Opa.Bool(false),
						Protocols:     Opa.StringSlice("http", "https"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &stateBuilder{
				currentState: tt.fields.currentState,
			}
			b.rawState = &utils.OpaRawState{}
			d, _ := utils.GetOpaDefaulter()
			b.defaulter = d
			b.intermediate, _ = state.NewOpaState()
			if err := b.ingestPolicy(tt.args.route); (err != nil) != tt.wantErr {
				t.Errorf("stateBuilder.ingestPlugins() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(tt.wantState, b.rawState)
		})
	}
}

func Test_stateBuilder_ingestTargets(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		currentState *state.OpaState
	}
	type args struct {
		targets []Opa.Target
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantState *utils.OpaRawState
	}{
		{
			name: "generates ID for a non-existing target",
			fields: fields{
				currentState: emptyState(),
			},
			args: args{
				targets: []Opa.Target{
					{
						Target: Opa.String("foo"),
						Upstream: &Opa.Upstream{
							ID: Opa.String("952ddf37-e815-40b6-b119-5379a3b1f7be"),
						},
					},
				},
			},
			wantErr: false,
			wantState: &utils.OpaRawState{
				Targets: []*Opa.Target{
					{
						ID:     Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Target: Opa.String("foo"),
						Weight: Opa.Int(100),
						Upstream: &Opa.Upstream{
							ID: Opa.String("952ddf37-e815-40b6-b119-5379a3b1f7be"),
						},
					},
				},
			},
		},
		{
			name: "matches up IDs of Targets correctly",
			fields: fields{
				currentState: existingTargetsState(),
			},
			args: args{
				targets: []Opa.Target{
					{
						Target: Opa.String("bar"),
						Upstream: &Opa.Upstream{
							ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
						},
					},
					{
						Target: Opa.String("foo"),
						Upstream: &Opa.Upstream{
							ID: Opa.String("700bc504-b2b1-4abd-bd38-cec92779659e"),
						},
					},
				},
			},
			wantErr: false,
			wantState: &utils.OpaRawState{
				Targets: []*Opa.Target{
					{
						ID:     Opa.String("f7e64af5-e438-4a9b-8ff8-ec6f5f06dccb"),
						Target: Opa.String("bar"),
						Weight: Opa.Int(100),
						Upstream: &Opa.Upstream{
							ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
						},
					},
					{
						ID:     Opa.String("53ce0a9c-d518-40ee-b8ab-1ee83a20d382"),
						Target: Opa.String("foo"),
						Weight: Opa.Int(100),
						Upstream: &Opa.Upstream{
							ID: Opa.String("700bc504-b2b1-4abd-bd38-cec92779659e"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &stateBuilder{
				currentState: tt.fields.currentState,
			}
			b.rawState = &utils.OpaRawState{}
			d, _ := utils.GetOpaDefaulter()
			b.defaulter = d
			if err := b.ingestTargets(tt.args.targets); (err != nil) != tt.wantErr {
				t.Errorf("stateBuilder.ingestPlugins() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(tt.wantState, b.rawState)
		})
	}
}

func Test_stateBuilder_ingestPlugins(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		currentState *state.OpaState
	}
	type args struct {
		plugins []FPlugin
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantState *utils.OpaRawState
	}{
		{
			name: "generates ID for a non-existing plugin",
			fields: fields{
				currentState: emptyState(),
			},
			args: args{
				plugins: []FPlugin{
					{
						Plugin: Opa.Plugin{
							Name: Opa.String("foo"),
						},
					},
				},
			},
			wantErr: false,
			wantState: &utils.OpaRawState{
				Plugins: []*Opa.Plugin{
					{
						ID:     Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Name:   Opa.String("foo"),
						Config: Opa.Configuration{},
					},
				},
			},
		},
		{
			name: "matches up IDs of plugins correctly",
			fields: fields{
				currentState: existingPluginState(),
			},
			args: args{
				plugins: []FPlugin{
					{
						Plugin: Opa.Plugin{
							Name: Opa.String("foo"),
						},
					},
					{
						Plugin: Opa.Plugin{
							Name: Opa.String("bar"),
							Consumer: &Opa.Consumer{
								ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
							},
						},
					},
					{
						Plugin: Opa.Plugin{
							Name: Opa.String("foo"),
							Consumer: &Opa.Consumer{
								ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
							},
							Route: &Opa.Route{
								ID: Opa.String("700bc504-b2b1-4abd-bd38-cec92779659e"),
							},
						},
					},
				},
			},
			wantErr: false,
			wantState: &utils.OpaRawState{
				Plugins: []*Opa.Plugin{
					{
						ID:     Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Name:   Opa.String("foo"),
						Config: Opa.Configuration{},
					},
					{
						ID:   Opa.String("f7e64af5-e438-4a9b-8ff8-ec6f5f06dccb"),
						Name: Opa.String("bar"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
						},
						Config: Opa.Configuration{},
					},
					{
						ID:   Opa.String("53ce0a9c-d518-40ee-b8ab-1ee83a20d382"),
						Name: Opa.String("foo"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("f77ca8c7-581d-45a4-a42c-c003234228e1"),
						},
						Route: &Opa.Route{
							ID: Opa.String("700bc504-b2b1-4abd-bd38-cec92779659e"),
						},
						Config: Opa.Configuration{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &stateBuilder{
				currentState: tt.fields.currentState,
			}
			b.rawState = &utils.OpaRawState{}
			if err := b.ingestPlugins(tt.args.plugins); (err != nil) != tt.wantErr {
				t.Errorf("stateBuilder.ingestPlugins() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(tt.wantState, b.rawState)
		})
	}
}

func Test_pluginRelations(t *testing.T) {
	type args struct {
		plugin *Opa.Plugin
	}
	tests := []struct {
		name    string
		args    args
		wantCID string
		wantRID string
		wantSID string
	}{
		{
			args: args{
				plugin: &Opa.Plugin{
					Name: Opa.String("foo"),
				},
			},
			wantCID: "",
			wantRID: "",
			wantSID: "",
		},
		{
			args: args{
				plugin: &Opa.Plugin{
					Name: Opa.String("foo"),
					Consumer: &Opa.Consumer{
						ID: Opa.String("cID"),
					},
					Route: &Opa.Route{
						ID: Opa.String("rID"),
					},
					Service: &Opa.Service{
						ID: Opa.String("sID"),
					},
				},
			},
			wantCID: "cID",
			wantRID: "rID",
			wantSID: "sID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCID, gotRID, gotSID := pluginRelations(tt.args.plugin)
			if gotCID != tt.wantCID {
				t.Errorf("pluginRelations() gotCID = %v, want %v", gotCID, tt.wantCID)
			}
			if gotRID != tt.wantRID {
				t.Errorf("pluginRelations() gotRID = %v, want %v", gotRID, tt.wantRID)
			}
			if gotSID != tt.wantSID {
				t.Errorf("pluginRelations() gotSID = %v, want %v", gotSID, tt.wantSID)
			}
		})
	}
}

func Test_stateBuilder_consumers(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		currentState  *state.OpaState
		targetContent *Content
	}
	tests := []struct {
		name   string
		fields fields
		want   *utils.OpaRawState
	}{
		{
			name: "generates ID for a non-existing consumer",
			fields: fields{
				targetContent: &Content{
					Consumers: []FConsumer{
						{
							Consumer: Opa.Consumer{
								Username: Opa.String("foo"),
							},
						},
					},
					Info: &Info{
						SelectorTags: []string{"tag1"},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				Consumers: []*Opa.Consumer{
					{
						ID:       Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Username: Opa.String("foo"),
						Tags:     Opa.StringSlice("tag1"),
					},
				},
			},
		},
		{
			name: "generates ID for a non-existing credential",
			fields: fields{
				targetContent: &Content{
					Consumers: []FConsumer{
						{
							Consumer: Opa.Consumer{
								Username: Opa.String("foo"),
							},
							KeyAuths: []*Opa.KeyAuth{
								{
									Key: Opa.String("foo-key"),
								},
							},
							BasicAuths: []*Opa.BasicAuth{
								{
									Username: Opa.String("basic-username"),
									Password: Opa.String("basic-password"),
								},
							},
							HMACAuths: []*Opa.HMACAuth{
								{
									Username: Opa.String("hmac-username"),
									Secret:   Opa.String("hmac-secret"),
								},
							},
							JWTAuths: []*Opa.JWTAuth{
								{
									Key:    Opa.String("jwt-key"),
									Secret: Opa.String("jwt-secret"),
								},
							},
							Oauth2Creds: []*Opa.Oauth2Credential{
								{
									ClientID: Opa.String("oauth2-clientid"),
									Name:     Opa.String("oauth2-name"),
								},
							},
							ACLGroups: []*Opa.ACLGroup{
								{
									Group: Opa.String("foo-group"),
								},
							},
						},
					},
					Info: &Info{
						SelectorTags: []string{"tag1"},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				Consumers: []*Opa.Consumer{
					{
						ID:       Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						Username: Opa.String("foo"),
						Tags:     Opa.StringSlice("tag1"),
					},
				},
				KeyAuths: []*Opa.KeyAuth{
					{
						ID:  Opa.String("dfd79b4d-7642-4b61-ba0c-9f9f0d3ba55b"),
						Key: Opa.String("foo-key"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				BasicAuths: []*Opa.BasicAuth{
					{
						ID:       Opa.String("0cc0d614-4c88-4535-841a-cbe0709b0758"),
						Username: Opa.String("basic-username"),
						Password: Opa.String("basic-password"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				HMACAuths: []*Opa.HMACAuth{
					{
						ID:       Opa.String("083f61d3-75bc-42b4-9df4-f91929e18fda"),
						Username: Opa.String("hmac-username"),
						Secret:   Opa.String("hmac-secret"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				JWTAuths: []*Opa.JWTAuth{
					{
						ID:     Opa.String("9e6f82e5-4e74-4e81-a79e-4bbd6fe34cdc"),
						Key:    Opa.String("jwt-key"),
						Secret: Opa.String("jwt-secret"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				Oauth2Creds: []*Opa.Oauth2Credential{
					{
						ID:       Opa.String("ba843ee8-d63e-4c4f-be1c-ebea546d8fac"),
						ClientID: Opa.String("oauth2-clientid"),
						Name:     Opa.String("oauth2-name"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				ACLGroups: []*Opa.ACLGroup{
					{
						ID:    Opa.String("13dd1aac-04ce-4ea2-877c-5579cfa2c78e"),
						Group: Opa.String("foo-group"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
			},
		},
		{
			name: "matches ID of an existing consumer",
			fields: fields{
				targetContent: &Content{
					Consumers: []FConsumer{
						{
							Consumer: Opa.Consumer{
								Username: Opa.String("foo"),
							},
						},
					},
				},
				currentState: existingConsumerCredState(),
			},
			want: &utils.OpaRawState{
				Consumers: []*Opa.Consumer{
					{
						ID:       Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Username: Opa.String("foo"),
					},
				},
			},
		},
		{
			name: "matches ID of an existing credential",
			fields: fields{
				targetContent: &Content{
					Consumers: []FConsumer{
						{
							Consumer: Opa.Consumer{
								Username: Opa.String("foo"),
							},
							KeyAuths: []*Opa.KeyAuth{
								{
									Key: Opa.String("foo-apikey"),
								},
							},
							BasicAuths: []*Opa.BasicAuth{
								{
									Username: Opa.String("basic-username"),
									Password: Opa.String("basic-password"),
								},
							},
							HMACAuths: []*Opa.HMACAuth{
								{
									Username: Opa.String("hmac-username"),
									Secret:   Opa.String("hmac-secret"),
								},
							},
							JWTAuths: []*Opa.JWTAuth{
								{
									Key:    Opa.String("jwt-key"),
									Secret: Opa.String("jwt-secret"),
								},
							},
							Oauth2Creds: []*Opa.Oauth2Credential{
								{
									ClientID: Opa.String("oauth2-clientid"),
									Name:     Opa.String("oauth2-name"),
								},
							},
							ACLGroups: []*Opa.ACLGroup{
								{
									Group: Opa.String("foo-group"),
								},
							},
						},
					},
					Info: &Info{
						SelectorTags: []string{"tag1"},
					},
				},
				currentState: existingConsumerCredState(),
			},
			want: &utils.OpaRawState{
				Consumers: []*Opa.Consumer{
					{
						ID:       Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Username: Opa.String("foo"),
						Tags:     Opa.StringSlice("tag1"),
					},
				},
				KeyAuths: []*Opa.KeyAuth{
					{
						ID:  Opa.String("5f1ef1ea-a2a5-4a1b-adbb-b0d3434013e5"),
						Key: Opa.String("foo-apikey"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				BasicAuths: []*Opa.BasicAuth{
					{
						ID:       Opa.String("92f4c849-960b-43af-aad3-f307051408d3"),
						Username: Opa.String("basic-username"),
						Password: Opa.String("basic-password"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				HMACAuths: []*Opa.HMACAuth{
					{
						ID:       Opa.String("e5d81b73-bf9e-42b0-9d68-30a1d791b9c9"),
						Username: Opa.String("hmac-username"),
						Secret:   Opa.String("hmac-secret"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				JWTAuths: []*Opa.JWTAuth{
					{
						ID:     Opa.String("917b9402-1be0-49d2-b482-ca4dccc2054e"),
						Key:    Opa.String("jwt-key"),
						Secret: Opa.String("jwt-secret"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				Oauth2Creds: []*Opa.Oauth2Credential{
					{
						ID:       Opa.String("4eef5285-3d6a-4f6b-b659-8957a940e2ca"),
						ClientID: Opa.String("oauth2-clientid"),
						Name:     Opa.String("oauth2-name"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				ACLGroups: []*Opa.ACLGroup{
					{
						ID:    Opa.String("b7c9352a-775a-4ba5-9869-98e926a3e6cb"),
						Group: Opa.String("foo-group"),
						Consumer: &Opa.Consumer{
							ID: Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						},
						Tags: Opa.StringSlice("tag1"),
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

func Test_stateBuilder_certificates(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		currentState  *state.OpaState
		targetContent *Content
	}
	tests := []struct {
		name   string
		fields fields
		want   *utils.OpaRawState
	}{
		{
			name: "generates ID for a non-existing certificate",
			fields: fields{
				targetContent: &Content{
					Certificates: []FCertificate{
						{
							Certificate: Opa.Certificate{
								Cert: Opa.String("foo"),
								Key:  Opa.String("bar"),
							},
						},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				Certificates: []*Opa.Certificate{
					{
						ID:   Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Cert: Opa.String("foo"),
						Key:  Opa.String("bar"),
					},
				},
			},
		},
		{
			name: "matches ID of an existing certificate",
			fields: fields{
				targetContent: &Content{
					Certificates: []FCertificate{
						{
							Certificate: Opa.Certificate{
								Cert: Opa.String("foo"),
								Key:  Opa.String("bar"),
							},
						},
					},
				},
				currentState: existingCertificateState(),
			},
			want: &utils.OpaRawState{
				Certificates: []*Opa.Certificate{
					{
						ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Cert: Opa.String("foo"),
						Key:  Opa.String("bar"),
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

func Test_stateBuilder_caCertificates(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(42)
	type fields struct {
		currentState  *state.OpaState
		targetContent *Content
	}
	tests := []struct {
		name   string
		fields fields
		want   *utils.OpaRawState
	}{
		{
			name: "generates ID for a non-existing CACertificate",
			fields: fields{
				targetContent: &Content{
					CACertificates: []FCACertificate{
						{
							CACertificate: Opa.CACertificate{
								Cert: Opa.String("foo"),
							},
						},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				CACertificates: []*Opa.CACertificate{
					{
						ID:   Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Cert: Opa.String("foo"),
					},
				},
			},
		},
		{
			name: "matches ID of an existing CACertificate",
			fields: fields{
				targetContent: &Content{
					CACertificates: []FCACertificate{
						{
							CACertificate: Opa.CACertificate{
								Cert: Opa.String("foo"),
							},
						},
					},
				},
				currentState: existingCACertificateState(),
			},
			want: &utils.OpaRawState{
				CACertificates: []*Opa.CACertificate{
					{
						ID:   Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Cert: Opa.String("foo"),
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

func Test_stateBuilder_upstream(t *testing.T) {
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
			name: "process a non-existent upstream",
			fields: fields{
				targetContent: &Content{
					Info: &Info{
						SelectorTags: []string{"tag1"},
					},
					Upstreams: []FUpstream{
						{
							Upstream: Opa.Upstream{
								Name:  Opa.String("foo"),
								Slots: Opa.Int(42),
							},
						},
					},
				},
				currentState: existingServiceState(),
			},
			want: &utils.OpaRawState{
				Upstreams: []*Opa.Upstream{
					{
						ID:    Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Name:  Opa.String("foo"),
						Slots: Opa.Int(42),
						Healthchecks: &Opa.Healthcheck{
							Active: &Opa.ActiveHealthcheck{
								Concurrency: Opa.Int(10),
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 302},
									Interval:     Opa.Int(0),
									Successes:    Opa.Int(0),
								},
								HTTPPath:               Opa.String("/"),
								HTTPSVerifyCertificate: Opa.Bool(true),
								Type:                   Opa.String("http"),
								Timeout:                Opa.Int(1),
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									Interval:     Opa.Int(0),
									HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
								},
							},
							Passive: &Opa.PassiveHealthcheck{
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 201, 202, 203, 204, 205,
										206, 207, 208, 226, 300, 301, 302, 303, 304, 305,
										306, 307, 308},
									Successes: Opa.Int(0),
								},
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									HTTPStatuses: []int{429, 500, 503},
								},
							},
						},
						HashOn:           Opa.String("none"),
						HashFallback:     Opa.String("none"),
						HashOnCookiePath: Opa.String("/"),
						Tags:             Opa.StringSlice("tag1"),
					},
				},
			},
		},
		{
			name: "matches ID of an existing policy",
			fields: fields{
				targetContent: &Content{
					Upstreams: []FUpstream{
						{
							Upstream: Opa.Upstream{
								Name: Opa.String("foo"),
							},
						},
					},
				},
				currentState: existingUpstreamState(),
			},
			want: &utils.OpaRawState{
				Upstreams: []*Opa.Upstream{
					{
						ID:    Opa.String("4bfcb11f-c962-4817-83e5-9433cf20b663"),
						Name:  Opa.String("foo"),
						Slots: Opa.Int(10000),
						Healthchecks: &Opa.Healthcheck{
							Active: &Opa.ActiveHealthcheck{
								Concurrency: Opa.Int(10),
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 302},
									Interval:     Opa.Int(0),
									Successes:    Opa.Int(0),
								},
								HTTPPath:               Opa.String("/"),
								HTTPSVerifyCertificate: Opa.Bool(true),
								Type:                   Opa.String("http"),
								Timeout:                Opa.Int(1),
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									Interval:     Opa.Int(0),
									HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
								},
							},
							Passive: &Opa.PassiveHealthcheck{
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 201, 202, 203, 204, 205,
										206, 207, 208, 226, 300, 301, 302, 303, 304, 305,
										306, 307, 308},
									Successes: Opa.Int(0),
								},
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									HTTPStatuses: []int{429, 500, 503},
								},
							},
						},
						HashOn:           Opa.String("none"),
						HashFallback:     Opa.String("none"),
						HashOnCookiePath: Opa.String("/"),
					},
				},
			},
		},
		{
			name: "multiple upstreams are handled correctly",
			fields: fields{
				targetContent: &Content{
					Upstreams: []FUpstream{
						{
							Upstream: Opa.Upstream{
								Name: Opa.String("foo"),
							},
						},
						{
							Upstream: Opa.Upstream{
								Name: Opa.String("bar"),
							},
						},
					},
				},
				currentState: emptyState(),
			},
			want: &utils.OpaRawState{
				Upstreams: []*Opa.Upstream{
					{
						ID:    Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						Name:  Opa.String("foo"),
						Slots: Opa.Int(10000),
						Healthchecks: &Opa.Healthcheck{
							Active: &Opa.ActiveHealthcheck{
								Concurrency: Opa.Int(10),
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 302},
									Interval:     Opa.Int(0),
									Successes:    Opa.Int(0),
								},
								HTTPPath:               Opa.String("/"),
								HTTPSVerifyCertificate: Opa.Bool(true),
								Type:                   Opa.String("http"),
								Timeout:                Opa.Int(1),
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									Interval:     Opa.Int(0),
									HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
								},
							},
							Passive: &Opa.PassiveHealthcheck{
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 201, 202, 203, 204, 205,
										206, 207, 208, 226, 300, 301, 302, 303, 304, 305,
										306, 307, 308},
									Successes: Opa.Int(0),
								},
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									HTTPStatuses: []int{429, 500, 503},
								},
							},
						},
						HashOn:           Opa.String("none"),
						HashFallback:     Opa.String("none"),
						HashOnCookiePath: Opa.String("/"),
					},
					{
						ID:    Opa.String("dfd79b4d-7642-4b61-ba0c-9f9f0d3ba55b"),
						Name:  Opa.String("bar"),
						Slots: Opa.Int(10000),
						Healthchecks: &Opa.Healthcheck{
							Active: &Opa.ActiveHealthcheck{
								Concurrency: Opa.Int(10),
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 302},
									Interval:     Opa.Int(0),
									Successes:    Opa.Int(0),
								},
								HTTPPath:               Opa.String("/"),
								HTTPSVerifyCertificate: Opa.Bool(true),
								Type:                   Opa.String("http"),
								Timeout:                Opa.Int(1),
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									Interval:     Opa.Int(0),
									HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
								},
							},
							Passive: &Opa.PassiveHealthcheck{
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 201, 202, 203, 204, 205,
										206, 207, 208, 226, 300, 301, 302, 303, 304, 305,
										306, 307, 308},
									Successes: Opa.Int(0),
								},
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									HTTPStatuses: []int{429, 500, 503},
								},
							},
						},
						HashOn:           Opa.String("none"),
						HashFallback:     Opa.String("none"),
						HashOnCookiePath: Opa.String("/"),
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

func Test_stateBuilder(t *testing.T) {
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
			name: "end to end test with all entities",
			fields: fields{
				targetContent: &Content{
					Info: &Info{
						SelectorTags: []string{"tag1"},
					},
					Services: []FPolicy{
						{
							Service: Opa.Service{
								Name: Opa.String("foo-policy"),
							},
							Routes: []*FRoute{
								{
									Route: Opa.Route{
										Name: Opa.String("foo-route1"),
									},
								},
								{
									Route: Opa.Route{
										ID:   Opa.String("d125e79a-297c-414b-bc00-ad3a87be6c2b"),
										Name: Opa.String("foo-route2"),
									},
								},
							},
						},
						{
							Service: Opa.Service{
								Name: Opa.String("bar-policy"),
							},
							Routes: []*FRoute{
								{
									Route: Opa.Route{
										Name: Opa.String("bar-route1"),
									},
								},
								{
									Route: Opa.Route{
										Name: Opa.String("bar-route2"),
									},
								},
							},
						},
					},
					Upstreams: []FUpstream{
						{
							Upstream: Opa.Upstream{
								Name:  Opa.String("foo"),
								Slots: Opa.Int(42),
							},
						},
					},
				},
				currentState: existingServiceState(),
			},
			want: &utils.OpaRawState{
				Services: []*Opa.Service{
					{
						ID:             Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						Name:           Opa.String("foo-policy"),
						Port:           Opa.Int(80),
						Retries:        Opa.Int(5),
						Protocol:       Opa.String("http"),
						ConnectTimeout: Opa.Int(60000),
						WriteTimeout:   Opa.Int(60000),
						ReadTimeout:    Opa.Int(60000),
						Tags:           Opa.StringSlice("tag1"),
					},
					{
						ID:             Opa.String("dfd79b4d-7642-4b61-ba0c-9f9f0d3ba55b"),
						Name:           Opa.String("bar-policy"),
						Port:           Opa.Int(80),
						Retries:        Opa.Int(5),
						Protocol:       Opa.String("http"),
						ConnectTimeout: Opa.Int(60000),
						WriteTimeout:   Opa.Int(60000),
						ReadTimeout:    Opa.Int(60000),
						Tags:           Opa.StringSlice("tag1"),
					},
				},
				Routes: []*Opa.Route{
					{
						ID:            Opa.String("5b1484f2-5209-49d9-b43e-92ba09dd9d52"),
						Name:          Opa.String("foo-route1"),
						PreserveHost:  Opa.Bool(false),
						RegexPriority: Opa.Int(0),
						StripPath:     Opa.Bool(false),
						Protocols:     Opa.StringSlice("http", "https"),
						Service: &Opa.Service{
							ID: Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
					{
						ID:            Opa.String("d125e79a-297c-414b-bc00-ad3a87be6c2b"),
						Name:          Opa.String("foo-route2"),
						PreserveHost:  Opa.Bool(false),
						RegexPriority: Opa.Int(0),
						StripPath:     Opa.Bool(false),
						Protocols:     Opa.StringSlice("http", "https"),
						Service: &Opa.Service{
							ID: Opa.String("538c7f96-b164-4f1b-97bb-9f4bb472e89f"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
					{
						ID:            Opa.String("0cc0d614-4c88-4535-841a-cbe0709b0758"),
						Name:          Opa.String("bar-route1"),
						PreserveHost:  Opa.Bool(false),
						RegexPriority: Opa.Int(0),
						StripPath:     Opa.Bool(false),
						Protocols:     Opa.StringSlice("http", "https"),
						Service: &Opa.Service{
							ID: Opa.String("dfd79b4d-7642-4b61-ba0c-9f9f0d3ba55b"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
					{
						ID:            Opa.String("083f61d3-75bc-42b4-9df4-f91929e18fda"),
						Name:          Opa.String("bar-route2"),
						PreserveHost:  Opa.Bool(false),
						RegexPriority: Opa.Int(0),
						StripPath:     Opa.Bool(false),
						Protocols:     Opa.StringSlice("http", "https"),
						Service: &Opa.Service{
							ID: Opa.String("dfd79b4d-7642-4b61-ba0c-9f9f0d3ba55b"),
						},
						Tags: Opa.StringSlice("tag1"),
					},
				},
				Upstreams: []*Opa.Upstream{
					{
						ID:    Opa.String("9e6f82e5-4e74-4e81-a79e-4bbd6fe34cdc"),
						Name:  Opa.String("foo"),
						Slots: Opa.Int(42),
						Healthchecks: &Opa.Healthcheck{
							Active: &Opa.ActiveHealthcheck{
								Concurrency: Opa.Int(10),
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 302},
									Interval:     Opa.Int(0),
									Successes:    Opa.Int(0),
								},
								HTTPPath:               Opa.String("/"),
								HTTPSVerifyCertificate: Opa.Bool(true),
								Type:                   Opa.String("http"),
								Timeout:                Opa.Int(1),
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									Interval:     Opa.Int(0),
									HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
								},
							},
							Passive: &Opa.PassiveHealthcheck{
								Healthy: &Opa.Healthy{
									HTTPStatuses: []int{200, 201, 202, 203, 204, 205,
										206, 207, 208, 226, 300, 301, 302, 303, 304, 305,
										306, 307, 308},
									Successes: Opa.Int(0),
								},
								Unhealthy: &Opa.Unhealthy{
									HTTPFailures: Opa.Int(0),
									TCPFailures:  Opa.Int(0),
									Timeouts:     Opa.Int(0),
									HTTPStatuses: []int{429, 500, 503},
								},
							},
						},
						HashOn:           Opa.String("none"),
						HashFallback:     Opa.String("none"),
						HashOnCookiePath: Opa.String("/"),
						Tags:             Opa.StringSlice("tag1"),
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

func Test_stateBuilder_fillPluginConfig(t *testing.T) {
	type fields struct {
		targetContent *Content
	}
	type args struct {
		plugin *FPlugin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		result  FPlugin
	}{
		{
			name:    "nil arg throws an error",
			wantErr: true,
		},
		{
			name: "no _plugin_config throws an error",
			fields: fields{
				targetContent: &Content{},
			},
			args: args{
				plugin: &FPlugin{
					ConfigSource: Opa.String("foo"),
				},
			},
			wantErr: true,
		},
		{
			name: "no _plugin_config throws an error",
			fields: fields{
				targetContent: &Content{
					PluginConfigs: map[string]Opa.Configuration{
						"foo": {
							"k2":  "v3",
							"k3:": "v3",
						},
					},
				},
			},
			args: args{
				plugin: &FPlugin{
					ConfigSource: Opa.String("foo"),
					Plugin: Opa.Plugin{
						Config: Opa.Configuration{
							"k1": "v1",
							"k2": "v2",
						},
					},
				},
			},
			result: FPlugin{
				ConfigSource: Opa.String("foo"),
				Plugin: Opa.Plugin{
					Config: Opa.Configuration{
						"k1":  "v1",
						"k2":  "v2",
						"k3:": "v3",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &stateBuilder{
				targetContent: tt.fields.targetContent,
			}
			if err := b.fillPluginConfig(tt.args.plugin); (err != nil) != tt.wantErr {
				t.Errorf("stateBuilder.fillPluginConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(tt.result, tt.args.plugin) {
				assert.Equal(t, tt.result, *tt.args.plugin)
			}
		})
	}
}
