package openapi

import "encoding/json"

// path.go contains
//
// PathItemObject
// OperationObject
// ParameterObject

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#path-item-object
type Path struct {
	// TODO: how to deal with Ref, we could inline everything when generate though ....
	//Ref         string      `json:"$ref" yaml:"$ref"`
	Summary     string      `json:"summary" yaml:"summary"`
	Description string      `json:"description" yaml:"description"`
	Servers     []Server    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Get         *Operation  `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation  `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation  `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation  `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *Operation  `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation  `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *Operation  `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *Operation  `json:"trace,omitempty" yaml:"trace,omitempty"`
}

// Operations describes a single API call <HTTP verb> <path>, i.e. GET /pets
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#operation-object
type Operation struct {
	Tags         []string            `json:"tags" yaml:"tags"`
	Summary      string              `json:"summary" yaml:"summary"`
	Description  string              `json:"description" yaml:"description"`
	ExternalDocs *ExternalDoc        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationId  string              `json:"operationId" yaml:"operationId"`
	Parameters   []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBody        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]Response `json:"responses" yaml:"responses"`
	// TODO: callbacks
	Deprecated bool      `json:"deprecated" yaml:"deprecated"`
	Security   *Security `json:"security,omitempty" yaml:"security,omitempty"`
	Servers    []Server  `json:"servers,omitempty" yaml:"servers,omitempty"`
}

type ParameterOrRef struct {
	Ref       Reference
	Parameter Parameter
}

func (r ParameterOrRef) MarshalJSON() ([]byte, error) {
	// use reference
	if r.Ref.Ref != "" {
		return json.Marshal(r.Ref)
	}
	return json.Marshal(r.Parameter)
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#parameter-object
// {
//  "name": "token",
//  "in": "header",
//  "description": "token to be passed as a header",
//  "required": true,
//  "schema": {
//    "type": "array",
//    "items": {
//      "type": "integer",
//      "format": "int64"
//    }
//  },
//  "style": "simple"
// }
type Parameter struct {
	Name            string       `json:"name" yaml:"name"`
	In              string       `json:"in" yaml:"in"`
	Description     string       `json:"description" yaml:"description"`
	Schema          *SchemaOrRef `json:"schema" yaml:"schema"`
	Style           string       `json:"style" yaml:"style"`
	Required        bool         `json:"required" yaml:"required"`
	Deprecated      bool         `json:"deprecated" yaml:"deprecated"`
	AllowEmptyValue bool         `json:"allowEmptyValue" yaml:"allowEmptyValue"`
}

type RequestBodyOrRef struct {
	Ref         Reference
	RequestBody RequestBody
}

func (r RequestBodyOrRef) MarshalJSON() ([]byte, error) {
	// use reference
	if r.Ref.Ref != "" {
		return json.Marshal(r.Ref)
	}
	return json.Marshal(r.RequestBody)
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#requestBodyObject
type RequestBody struct {
	Description string               `json:"description" yaml:"description"`
	Content     map[string]MediaType `json:"content" yaml:"content"`
	Required    bool                 `json:"required" yaml:"required"`
}

type ResponseOrRef struct {
	Ref      Reference
	Response Response
}

func (r ResponseOrRef) MarshalJSON() ([]byte, error) {
	// use reference
	if r.Ref.Ref != "" {
		return json.Marshal(r.Ref)
	}
	return json.Marshal(r.Response)
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#responseObject
type Response struct {
	Description string `json:"description" yaml:"description"`
	// TODO: headers, need to define header object https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#headerObject
	// TODO: might need to use pointer to indicate there is no response body
	Content map[string]MediaType `json:"content" yaml:"content"`
	// TODO: links
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#mediaTypeObject
type MediaType struct {
	Schema *SchemaOrRef `json:"schema" yaml:"schema"`
	// TODO: example, examples
}
