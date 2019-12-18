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
	if s.Name != nil {
		return *s.Name
	}
	return ""
}

type policy struct {
	CreatedAt *int    `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	ID        *string `json:"id,omitempty" yaml:"id,omitempty"`
	Name      *string `json:"name,omitempty" yaml:"name,omitempty"`
	UpdatedAt *int    `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

func copyToPolicy(fPolicy FPolicy) policy {
	s := policy{}
	s.ID = fPolicy.ID
	s.Name = fPolicy.Name

	return s
}

func copyFromPolicy(policy policy, fService *FPolicy) {
	fService.CreatedAt = policy.CreatedAt
	fService.ID = policy.ID
	fService.Name = policy.Name
	fService.UpdatedAt = policy.UpdatedAt
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

// Info contains meta-data of the file.
type Info struct {
	SelectorTags []string `json:"select_tags,omitempty" yaml:"select_tags,omitempty"`
}

//go:generate go run ./codegen/main.go

// Content represents a serialized opa state.
type Content struct {
	FormatVersion string `json:"_format_version,omitempty" yaml:"_format_version,omitempty"`
	Info          *Info  `json:"_info,omitempty" yaml:"_info,omitempty"`
	Workspace     string `json:"_workspace,omitempty" yaml:"_workspace,omitempty"`

	Policies []FPolicy `json:"policies,omitempty" yaml:",omitempty"`
}
