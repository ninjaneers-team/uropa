package opa

// Policy represents a Policy in Opa.
// +k8s:deepcopy-gen=true
type Policy struct {
	ID  *string `json:"id,omitempty" yaml:"id,omitempty"`
	Raw *string `json:"raw,omitempty" yaml:"raw,omitempty"`
}

// Configuration represents a config of a plugin in Opa.
type Configuration map[string]interface{}

// DeepCopyInto copies the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
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
