package openapi

// Info is meta about the API, following fields are required
// - title
// - version
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#info-object
// {
//  "title": "Sample Pet Store App",
//  "description": "This is a sample server for a pet store.",
//  "termsOfService": "http://example.com/terms/",
//  "contact": {
//    "name": "API Support",
//    "url": "http://www.example.com/support",
//    "email": "support@example.com"
//  },
//  "license": {
//    "name": "Apache 2.0",
//    "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
//  },
//  "version": "1.0.1"
// }
type Info struct {
	Title          string   `json:"title" yaml:"title"`
	Description    string   `json:"description" yaml:"description"`
	TermsOfService string   `json:"termsOfService" yaml:"termsOfService"`
	Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string   `json:"version" yaml:"version"`
}

// Contact is optional, it contains contact info of exposed API
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#contact-object
// {
//  "name": "API Support",
//  "url": "http://www.example.com/support",
//  "email": "support@example.com"
// }
type Contact struct {
	Name  string `json:"name" yaml:"name"`
	Url   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}

// License is license of the exposed API
// TODO: I don't know there are license for web API, is implementing a service using API violation of API license?
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#licenseObject
// {
//  "name": "Apache 2.0",
//  "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
// }
type License struct {
	Name string `json:"name" yaml:"name"`
	Url  string `json:"url" yaml:"url"`
}

// Server is endpoint of a remote server that can be used to test the API, it can be
// put into multi levels of the doc and the inner level overrides the outer levels
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#serverObject
// {
//  "servers": [
//    {
//      "url": "https://development.gigantic-server.com/v1",
//      "description": "Development server"
//    },
//    {
//      "url": "https://staging.gigantic-server.com/v1",
//      "description": "Staging server"
//    },
//    {
//      "url": "https://api.gigantic-server.com/v1",
//      "description": "Production server"
//    }
//  ]
// }
type Server struct {
	Url         string `json:"url" yaml:"url"`
	Description string `json:"description" yaml:"description"`
	// TODO: variables
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#tagObject
// {
//	"name": "pet",
//	"description": "Pets operations"
// }
type Tag struct {
	Name         string       `json:"name" yaml:"name"`
	Description  string       `json:"description" yaml:"description"`
	ExternalDocs *ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#externalDocumentationObject
// {
//  "description": "Find more info here",
//  "url": "https://example.com"
// }
type ExternalDoc struct {
	Description string `json:"description" yaml:"description"`
	Url         string `json:"url" yaml:"url"`
}
