package openapi

import (
	"encoding/json"
)

type SchemaOrRef struct {
	Ref    Reference
	Schema Schema
}

// MarshalJSON use reference if provided, otherwise it encode the inline schema
// it implements https://golang.org/pkg/encoding/json/#Marshaler interface
func (r *SchemaOrRef) MarshalJSON() ([]byte, error) {
	// use reference
	if r.Ref.Ref != "" {
		return json.Marshal(r.Ref)
	}
	return json.Marshal(r.Schema)
}

// it implements https://godoc.org/gopkg.in/yaml.v2#Marshaler interface
func (r *SchemaOrRef) MarshalYAML() (interface{}, error) {
	if r.Ref.Ref != "" {
		return r.Ref, nil
	}
	return r.Schema, nil
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schema-object
// TODO: only a small subset is included, should be enough to describe api without validation
// throw people to json schema doc makes the hard part in open api pretty easy for the writer
// TODO: ref https://github.com/googleapis/gnostic/blob/master/jsonschema/models.go
type Schema struct {
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Type   string `json:"type" yaml:"type"`
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// object
	Properties map[string]*SchemaOrRef `json:"properties,omitempty" yaml:"properties,omitempty"`
	Required   []string                `json:"required,omitempty" yaml:"required,omitempty"`

	// array
	// TODO: items can even be a slice of schema ... tuple is allowed?
	// NOTE: use pointer to avoid `invalid recursive type Schema`
	// https://stackoverflow.com/questions/8261058/invalid-recursive-type-in-a-struct-in-go
	Items *SchemaOrRef `json:"items,omitempty" yaml:"items,omitempty"` // NOTE: it need
	// TODO: validation
}
