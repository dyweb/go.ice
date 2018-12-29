# OpenAPI

## Develop

- use [swagger-ui](https://github.com/swagger-api/swagger-editor) to valid and visualize generated YAML/json

Known issues

- need to use pointer to avoid `invalid recursive type` because size of struct must be know at compile time, 
its fields can not contains the struct itself (indirect inclusion is also not possible)

## Reference

- https://github.com/googleapis/gnostic/blob/master/jsonschema/models.go contains the full json schema for `Schema Object`
  - gnostic is used to generate protobuf from swagger