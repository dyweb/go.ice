package openapi

import "encoding/json"

type SchemaOrRef struct {
	Ref    Reference
	Schema Schema
}

// MarshalJSON use reference if provided, otherwise it encode the inline schema
// it implements https://golang.org/pkg/encoding/json/#Marshaler interface
func (sr SchemaOrRef) MarshalJSON() ([]byte, error) {
	// use reference
	if sr.Ref.Ref != "" {
		return json.Marshal(sr.Ref)
	}
	return json.Marshal(sr.Schema)
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schema-object
// TODO: only a small subset is included, should be enough to describe api without validation
// throw people to json schema doc makes the hard part in open api pretty easy for the writer
// TODO: ref https://github.com/googleapis/gnostic/blob/master/jsonschema/models.go
type Schema struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`

	Type       string                 `json:"type" yaml:"type"`
	Format     string                 `json:"format" yaml:"format"`
	Properties map[string]SchemaOrRef `json:"properties" yaml:"properties"`

	// TODO: validation
}
