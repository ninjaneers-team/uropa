package opa

// Policy represents a Policy in Opa.
// +k8s:deepcopy-gen=true
type Policy struct {
	ID        *string   `json:"id,omitempty" yaml:"id,omitempty"`
	Name      *string   `json:"name,omitempty" yaml:"name,omitempty"`
	Tags      []*string `json:"tags,omitempty" yaml:"tags,omitempty"`
	Rego      *string   `json:"rego,omitempty" yaml:"rego,omitempty"`
	CreatedAt *int      `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt *int      `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

// Configuration represents a config of a plugin in Opa.
type Configuration map[string]interface{}

// DeepCopyInto copies the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy copies the receiver, creating a new Configuration.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}
