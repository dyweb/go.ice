// Package openapi defines structs for OpenAPI v3 schema
//
// The example specs snippets in the comments are directly copied from
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md
package openapi

// TODO: use pointer and omit empty for struct
// TODO: deal with ref

const Version = "3.0.2"

// Reference is used to references models defined in current API doc and external doc
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#referenceObject
// {
//	"$ref": "#/components/schemas/Pet"
// }
type Reference struct {
	Ref string `json:"$ref" yaml:"$ref"`
}

// Document is the full API doc
// - paths defines routes and their operations
// - components defines data models
// - tags is how you group the paths
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#openapi-object
type Document struct {
	Openapi      string              `json:"openapi" yaml:"openapi"`
	Info         Info                `json:"info" yaml:"info"`
	Servers      []Server            `json:"servers" yaml:"servers"`
	Paths        map[string]PathItem `json:"paths" yaml:"paths"`
	Components   *Components         `json:"components,omitempty" yaml:"components,omitempty"`
	Security     *Security           `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []Tag               `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *ExternalDoc        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// Components contains reusable objects that can be referenced
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#componentsObject
type Components struct {
	Schemas   map[string]SchemaOrRef   `json:"schemas" yaml:"schemas"`
	Responses map[string]ResponseOrRef `json:"responses" yaml:"responses"`
	// TODO: parameters
	// TODO: examples
	RequestBodies map[string]RequestBodyOrRef `json:"requestBodies" yaml:"requestBodies"`
	// TODO: headers
	// TODO: securitySchemes TODO: it is not securitySchema?
	// TODO: links
	// TODO: callbacks
}

// TODO: security
type Security struct {
}
