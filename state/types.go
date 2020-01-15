package state

import (
	"github.com/ninjaneers-team/uropa/opa"
	"reflect"
)

// entity abstracts out common fields in a credentials.
// TODO generalize for each and every entity.
type entity interface {
	// ID of the cred.
	GetID() string
	// ID2 is the second endpoint key.
	GetID2() string
	// Consumer returns consumer ID associated with the cred.
	GetConsumer() string
}

// ConsoleString contains methods to be used to print
// entity to console.
type ConsoleString interface {
	// Console returns a string to uniquely identify an
	// entity in human-readable form.
	// It should have the ID or endpoint key along-with
	// foreign references if they exist.
	// It will be used to communicate to the human user
	// that this entity is undergoing some change.
	Console() string
}

// Meta contains additional information for an entity
// type Meta struct {
// 	Global *bool   `json:"global,omitempty" yaml:"global,omitempty"`
// 	Kind   *string `json:"type,omitempty" yaml:"type,omitempty"`
// }

// Meta stores metadata for any entity.
type Meta struct {
	metaMap map[string]interface{}
}

func (m *Meta) initMeta() {
	if m.metaMap == nil {
		m.metaMap = make(map[string]interface{})
	}
}

// AddMeta adds key->obj metadata.
// It will override the old obj in key is already present.
func (m *Meta) AddMeta(key string, obj interface{}) {
	m.initMeta()
	m.metaMap[key] = obj
}

// GetMeta returns the obj previously added using AddMeta().
// It returns nil if key is not present.
func (m *Meta) GetMeta(key string) interface{} {
	m.initMeta()
	return m.metaMap[key]
}

// Service represents a service in Opa.
// It adds some helper methods along with Meta to the original Service object.
type Policy struct {
	opa.Policy `yaml:",inline"`
	Meta
}

// Identifier returns the endpoint key ID.
func (s1 *Policy) Identifier() string {
	return *s1.ID
}

// Console returns an entity's identity in a human
// readable string.
func (s1 *Policy) Console() string {
	return s1.Identifier()
}

// Equal returns true if s1 and s2 are equal.
func (s1 *Policy) Equal(s2 *Policy) bool {
	return reflect.DeepEqual(s1.Policy, s2.Policy)
}

// EqualWithOpts returns true if s1 and s2 are equal.
// If ignoreID is set to true, IDs will be ignored while comparison.
// If ignoreTS is set to true, timestamp fields will be ignored.
func (s1 *Policy) EqualWithOpts(s2 *Policy,
	ignoreID bool, ignoreTS bool) bool {
	s1Copy := s1.Policy.DeepCopy()
	s2Copy := s2.Policy.DeepCopy()

	if ignoreID {
		s1Copy.ID = nil
		s2Copy.ID = nil
	}
	return reflect.DeepEqual(s1Copy, s2Copy)
}
