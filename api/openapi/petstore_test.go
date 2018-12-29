package openapi_test

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/dyweb/go.ice/api/openapi"
)

// petstore_test tests if we can generate a valid petstore example and
// decode those written by human/generated from other tools

func TestPetStore_Simple(t *testing.T) {
	doc := openapi.Document{
		Openapi: openapi.Version,
		Info: openapi.Info{
			Version:     "0.0.1",
			Title:       "Simple Pet Store",
			Description: "I am *mark down*",
		},
		Servers: []openapi.Server{
			{
				Url:         "https://dyweb.com/petstore/v1",
				Description: "Testing server",
			},
		},
		Tags: []openapi.Tag{
			{
				Name:        "pet",
				Description: "Everything about your **Pet**",
			},
		},
		Paths: map[string]openapi.Path{
			"/pets": {
				Get: &openapi.Operation{
					Summary:     "List all pets",
					OperationId: "listPets",
					Tags: []string{
						"pet",
					},
					Responses: map[string]openapi.Response{
						"200": {
							Description: "A paged array of pets",
							Content: map[string]openapi.MediaType{
								"application/json": {
									Schema: &openapi.SchemaOrRef{
										// TODO: might need to simplify reference to just use a string
										Ref: openapi.Reference{
											Ref: "#/components/schemas/Pets",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Components: openapi.Components{
			Schemas: map[string]*openapi.SchemaOrRef{
				"Pet": {
					Schema: openapi.Schema{
						Type: "object",
						Required: []string{
							"id", "name",
						},
						Properties: map[string]*openapi.SchemaOrRef{
							"id": {
								Schema: openapi.Schema{
									Type:   "integer",
									Format: "int64",
								},
							},
							"name": {
								Schema: openapi.Schema{
									Type: "string",
								},
							},
						},
					},
				},
				"Pets": {
					Schema: openapi.Schema{
						Type: "array",
						Items: &openapi.SchemaOrRef{
							Ref: openapi.Reference{
								Ref: "#/components/schemas/Pet",
							},
						},
					},
				},
			},
		},
	}
	dumpYAML(t, doc)
	saveYAML(t, doc, "testdata/petstore-simple.yaml")
}

// TODO: put in gommon as dump?
func dumpYAML(t *testing.T, val interface{}) {
	if b, err := yaml.Marshal(val); err != nil {
		t.Fatal(err)
	} else {
		os.Stdout.Write(b)
	}
}

func saveYAML(t *testing.T, val interface{}, p string) {
	b, err := yaml.Marshal(val)
	if err != nil {
		t.Fatal(err)
		return
	}
	if err := ioutil.WriteFile(p, b, 0664); err != nil {
		t.Fatal(err)
		return
	}
}
