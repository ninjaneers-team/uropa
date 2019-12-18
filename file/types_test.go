package file

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

var (
	jsonString = `{                                                                           
  "name": "rate-limiting",                                                  
  "config": {                                                               
    "minute": 10
  },                                                                        
  "policy": "foo",                                                         
  "route": "bar",                                                         
  "consumer": "baz",                                                        
  "enabled": true,                                                          
  "run_on": "first",                                                        
  "protocols": [                                                            
    "http"
  ]                                                                         
}`
	yamlString = `
name: rate-limiting
config:
  minute: 10
policy: foo
consumer: baz
route: bar
enabled: true
run_on: first
protocols:
- http
`
)

func TestPluginUnmarshalYAML(t *testing.T) {
	var p FPlugin
	assert := assert.New(t)
	assert.Nil(yaml.Unmarshal([]byte(yamlString), &p))
	assert.Equal(Opa.Plugin{
		Name:      p.Name,
		Config:    p.Config,
		Enabled:   p.Enabled,
		RunOn:     p.RunOn,
		Protocols: p.Protocols,
		Service: &Opa.Service{
			ID: Opa.String("foo"),
		},
		Consumer: &Opa.Consumer{
			ID: Opa.String("baz"),
		},
		Route: &Opa.Route{
			ID: Opa.String("bar"),
		},
	}, p.Plugin)
}

func TestPluginUnmarshalJSON(t *testing.T) {
	var p FPlugin
	assert := assert.New(t)
	assert.Nil(json.Unmarshal([]byte(jsonString), &p))
	assert.Equal(Opa.Plugin{
		Name:      p.Name,
		Config:    p.Config,
		Enabled:   p.Enabled,
		RunOn:     p.RunOn,
		Protocols: p.Protocols,
		Service: &Opa.Service{
			ID: Opa.String("foo"),
		},
		Consumer: &Opa.Consumer{
			ID: Opa.String("baz"),
		},
		Route: &Opa.Route{
			ID: Opa.String("bar"),
		},
	}, p.Plugin)
}

func Test_copyToCert(t *testing.T) {
	type args struct {
		certificate FCertificate
	}
	tests := []struct {
		name string
		args args
		want cert
	}{
		{
			name: "basic sanity",
			args: args{
				certificate: FCertificate{
					Certificate: Opa.Certificate{
						Key:  Opa.String("key"),
						Cert: Opa.String("cert"),
						ID:   Opa.String("cert-id"),
						SNIs: Opa.StringSlice("0.example.com", "1.example.com"),
						Tags: Opa.StringSlice("tag1", "tag2"),
					},
				},
			},
			want: cert{
				Key:  Opa.String("key"),
				Cert: Opa.String("cert"),
				ID:   Opa.String("cert-id"),
				SNIs: []*sni{
					{Name: Opa.String("0.example.com")},
					{Name: Opa.String("1.example.com")},
				},
				Tags: Opa.StringSlice("tag1", "tag2"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyToCert(tt.args.certificate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyToCert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copyFromCert(t *testing.T) {
	type args struct {
		cert        cert
		certificate *FCertificate
	}
	tests := []struct {
		name string
		args args
		want *FCertificate
	}{
		{
			name: "basic sanity",
			args: args{
				cert: cert{
					Key:  Opa.String("key"),
					Cert: Opa.String("cert"),
					ID:   Opa.String("cert-id"),
					SNIs: []*sni{
						{Name: Opa.String("0.example.com")},
						{Name: Opa.String("1.example.com")},
					},
					Tags: Opa.StringSlice("tag1", "tag2"),
				},
				certificate: &FCertificate{},
			},
			want: &FCertificate{
				Certificate: Opa.Certificate{
					Key:  Opa.String("key"),
					Cert: Opa.String("cert"),
					ID:   Opa.String("cert-id"),
					SNIs: Opa.StringSlice("0.example.com", "1.example.com"),
					Tags: Opa.StringSlice("tag1", "tag2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copyFromCert(tt.args.cert, tt.args.certificate)

			if !reflect.DeepEqual(tt.args.certificate, tt.want) {
				t.Errorf("copyFromCert() = %v, want %v", tt.args.certificate, tt.want)
			}
		})
	}
}
