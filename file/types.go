package file

import (
	"encoding/json"

	"github.com/ninjaneers-team/uropa/opa"
)

// Format is a file format for opa's configuration.
type Format string

type id interface {
	id() string
}

const (
	// JSON is JSON file format.
	JSON = "JSON"
	// YAML if YAML file format.
	YAML = "YAML"
)

// FPolicy represents a opa Service and it's associated routes and plugins.
type FPolicy struct {
	opa.Policy
}

// id is used for sorting.
func (s FPolicy) id() string {
	if s.ID != nil {
		return *s.ID
	}
	if s.ID != nil {
		return *s.ID
	}
	return ""
}

type policy struct {
	ID  *string `json:"id,omitempty" yaml:"id,omitempty"`
	Raw *string `json:"raw,omitempty" yaml:"raw,omitempty"`
}

func copyToPolicy(fPolicy FPolicy) policy {
	s := policy{}
	s.ID = fPolicy.ID
	s.Raw = fPolicy.Raw

	return s
}

func copyFromPolicy(policy policy, fPolicy *FPolicy) {
	fPolicy.ID = policy.ID
	fPolicy.Raw = policy.Raw
}

// MarshalYAML is a custom marshal to handle
// SNI.
func (s FPolicy) MarshalYAML() (interface{}, error) {
	return copyToPolicy(s), nil
}

// UnmarshalYAML is a custom marshal method to handle
// foreign references.
func (s *FPolicy) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var service policy
	if err := unmarshal(&service); err != nil {
		return err
	}
	copyFromPolicy(service, s)
	return nil
}

// MarshalJSON is a custom marshal method to handle
// foreign references.
func (s FPolicy) MarshalJSON() ([]byte, error) {
	service := copyToPolicy(s)
	return json.Marshal(service)
}

// UnmarshalJSON is a custom marshal method to handle
// foreign references.
func (s *FPolicy) UnmarshalJSON(b []byte) error {
	var service policy
	err := json.Unmarshal(b, &service)
	if err != nil {
		return err
	}
	copyFromPolicy(service, s)
	return nil
}

//go:generate go run ./codegen/main.go

// Content represents a serialized opa state.
type Content struct {
	FormatVersion string    `json:"_format_version,omitempty" yaml:"_format_version,omitempty"`
	Policies      []FPolicy `json:"policies,omitempty" yaml:",omitempty"`
}
