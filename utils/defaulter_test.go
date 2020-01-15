package utils

import (
	"testing"

	"github.com/ninjaneers-team/uropa/opa"
	"github.com/stretchr/testify/assert"
)

func TestDefaulter(t *testing.T) {
	assert := assert.New(t)

	var d Defaulter

	assert.NotNil(d.Register(nil))
	assert.NotNil(d.Set(nil))

	assert.Panics(func() {
		d.MustSet(d)
	})

	type Foo struct {
		A string
		B []string
	}
	defaultFoo := &Foo{
		A: "defaultA",
		B: []string{"default1"},
	}
	assert.Nil(d.Register(defaultFoo))

	// sets a default
	var arg Foo
	assert.Nil(d.Set(&arg))
	assert.Equal("defaultA", arg.A)
	assert.Equal([]string{"default1"}, arg.B)

	// doesn't set a default
	arg1 := Foo{
		A: "non-default-value",
	}
	assert.Nil(d.Set(&arg1))
	assert.Equal("non-default-value", arg1.A)

	// errors on an unregistered type
	type Bar struct {
		A string
	}
	assert.NotNil(d.Set(&Bar{}))
	assert.Panics(func() {
		d.MustSet(&Bar{})
	})
}

func TestServiceSetTest(t *testing.T) {
	assert := assert.New(t)
	d, err := GetOpaDefaulter()
	assert.NotNil(d)
	assert.Nil(err)

	testCases := []struct {
		desc string
		arg  *Opa.Service
		want *Opa.Service
	}{
		{
			desc: "empty service",
			arg:  &Opa.Service{},
			want: &policyDefaults,
		},
		{
			desc: "timeout value value is not overridden",
			arg: &Opa.Service{
				WriteTimeout: Opa.Int(42),
			},
			want: &Opa.Service{
				Port:           Opa.Int(80),
				Retries:        Opa.Int(5),
				Protocol:       Opa.String("http"),
				ConnectTimeout: Opa.Int(60000),
				WriteTimeout:   Opa.Int(42),
				ReadTimeout:    Opa.Int(60000),
			},
		},
		{
			desc: "path value is not overridden",
			arg: &Opa.Service{
				Path: Opa.String("/foo"),
			},
			want: &Opa.Service{
				Port:           Opa.Int(80),
				Retries:        Opa.Int(5),
				Protocol:       Opa.String("http"),
				Path:           Opa.String("/foo"),
				ConnectTimeout: Opa.Int(60000),
				WriteTimeout:   Opa.Int(60000),
				ReadTimeout:    Opa.Int(60000),
			},
		},
		{
			desc: "Name is not reset",
			arg: &Opa.Service{
				Name: Opa.String("foo"),
				Host: Opa.String("example.com"),
				Path: Opa.String("/bar"),
			},
			want: &Opa.Service{
				Name:           Opa.String("foo"),
				Host:           Opa.String("example.com"),
				Port:           Opa.Int(80),
				Retries:        Opa.Int(5),
				Protocol:       Opa.String("http"),
				Path:           Opa.String("/bar"),
				ConnectTimeout: Opa.Int(60000),
				WriteTimeout:   Opa.Int(60000),
				ReadTimeout:    Opa.Int(60000),
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := d.Set(tC.arg)
			assert.Nil(err)
			assert.Equal(tC.want, tC.arg)
		})
	}
}

func TestRouteSetTest(t *testing.T) {
	assert := assert.New(t)
	d, err := GetOpaDefaulter()
	assert.NotNil(d)
	assert.Nil(err)

	testCases := []struct {
		desc string
		arg  *Opa.Route
		want *Opa.Route
	}{
		{
			desc: "empty route",
			arg:  &Opa.Route{},
			want: &routeDefaults,
		},
		{
			desc: "preserve host is not overridden",
			arg: &Opa.Route{
				PreserveHost: Opa.Bool(true),
			},
			want: &Opa.Route{
				PreserveHost:  Opa.Bool(true),
				RegexPriority: Opa.Int(0),
				StripPath:     Opa.Bool(false),
				Protocols:     Opa.StringSlice("http", "https"),
			},
		},
		{
			desc: "Protocols is not reset",
			arg: &Opa.Route{
				Protocols: Opa.StringSlice("http", "tls"),
			},
			want: &Opa.Route{
				PreserveHost:  Opa.Bool(false),
				RegexPriority: Opa.Int(0),
				StripPath:     Opa.Bool(false),
				Protocols:     Opa.StringSlice("http", "tls"),
			},
		},
		{
			desc: "non-default feilds is not reset",
			arg: &Opa.Route{
				Name:      Opa.String("foo"),
				Hosts:     Opa.StringSlice("1.example.com", "2.example.com"),
				Methods:   Opa.StringSlice("GET", "POST"),
				StripPath: Opa.Bool(false),
			},
			want: &Opa.Route{
				Name:          Opa.String("foo"),
				Hosts:         Opa.StringSlice("1.example.com", "2.example.com"),
				Methods:       Opa.StringSlice("GET", "POST"),
				PreserveHost:  Opa.Bool(false),
				RegexPriority: Opa.Int(0),
				StripPath:     Opa.Bool(false),
				Protocols:     Opa.StringSlice("http", "https"),
			},
		},
		{
			desc: "strip-path can be set to false",
			arg: &Opa.Route{
				StripPath: Opa.Bool(false),
			},
			want: &Opa.Route{
				PreserveHost:  Opa.Bool(false),
				RegexPriority: Opa.Int(0),
				StripPath:     Opa.Bool(false),
				Protocols:     Opa.StringSlice("http", "https"),
			},
		},
		{
			desc: "strip-path can be set to true",
			arg: &Opa.Route{
				StripPath: Opa.Bool(true),
			},
			want: &Opa.Route{
				PreserveHost:  Opa.Bool(false),
				RegexPriority: Opa.Int(0),
				StripPath:     Opa.Bool(true),
				Protocols:     Opa.StringSlice("http", "https"),
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := d.Set(tC.arg)
			assert.Nil(err)
			assert.Equal(tC.want, tC.arg)
		})
	}
}

func TestUpstreamSetTest(t *testing.T) {
	assert := assert.New(t)
	d, err := GetOpaDefaulter()
	assert.NotNil(d)
	assert.Nil(err)

	testCases := []struct {
		desc string
		arg  *Opa.Upstream
		want *Opa.Upstream
	}{
		{
			desc: "empty upstream",
			arg:  &Opa.Upstream{},
			want: &upstreamDefaults,
		},
		{
			desc: "Healthchecks.Active.Healthy.HTTPStatuses is not overridden",
			arg: &Opa.Upstream{
				Healthchecks: &Opa.Healthcheck{
					Active: &Opa.ActiveHealthcheck{
						Healthy: &Opa.Healthy{
							HTTPStatuses: []int{200},
						},
					},
				},
			},
			want: &Opa.Upstream{
				Slots: Opa.Int(10000),
				Healthchecks: &Opa.Healthcheck{
					Active: &Opa.ActiveHealthcheck{
						Concurrency: Opa.Int(10),
						Healthy: &Opa.Healthy{
							HTTPStatuses: []int{200},
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
							HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
							Interval:     Opa.Int(0),
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
		{
			desc: "Healthchecks.Active.Healthy.Timeout is not overridden",
			arg: &Opa.Upstream{
				Name: Opa.String("foo"),
				Healthchecks: &Opa.Healthcheck{
					Active: &Opa.ActiveHealthcheck{
						Healthy: &Opa.Healthy{
							Interval: Opa.Int(1),
						},
					},
				},
			},
			want: &Opa.Upstream{
				Name:  Opa.String("foo"),
				Slots: Opa.Int(10000),
				Healthchecks: &Opa.Healthcheck{
					Active: &Opa.ActiveHealthcheck{
						Concurrency: Opa.Int(10),
						Healthy: &Opa.Healthy{
							HTTPStatuses: []int{200, 302},
							Interval:     Opa.Int(1),
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
							HTTPStatuses: []int{429, 404, 500, 501, 502, 503, 504, 505},
							Interval:     Opa.Int(0),
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
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := d.Set(tC.arg)
			assert.Nil(err)
			assert.Equal(tC.want, tC.arg)
		})
	}
}
