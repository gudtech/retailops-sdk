package verify

import (
  // "fmt"

  "testing"
  schema "github.com/xeipuuv/gojsonschema"
)

func TestSchemaWithDefinitions(t *testing.T) {
  sampleSchema := `{
    "definitions": {
        "version_definition": {
            "type": [
                "integer"
            ],
            "description": "RetailOPS api action version",
            "example": 1
        }
    },
    "properties": {
      "version": {
        "$ref": "#/definitions/version_definition"
      }
    }
}`

  sampleDoc := `{
  "version": 1234
}`

  sampleSchemaLoader := schema.NewStringLoader(sampleSchema)
  sampleJsonLoader := schema.NewStringLoader(sampleDoc)

  result,err := schema.Validate(sampleSchemaLoader, sampleJsonLoader)
  if err != nil {
    t.Fatalf(err.Error())
    return
  }

  if result.Valid() {
  } else {
    t.Fatalf("failed")
  }

}