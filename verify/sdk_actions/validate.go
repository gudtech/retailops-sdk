package sdk_actions

import (
    "fmt"
     "encoding/json"
     gojsonschema "github.com/xeipuuv/gojsonschema"
     "os"
     "github.com/gudtech/scamp-go/scamp"
)

type V1FileLink struct {
    Description    string `json:"description"`
    Href           string `json:"href"`
    Method         string `json:"method"`
    Rel            string `json:"rel"`
    RequestSchema  *json.RawMessage `json:"schema"`
    ResponseSchema *json.RawMessage `json:"targetSchema"`
    Title          string `json:"title"`
}

type V1File struct {
    Schema      string `json:"$schema"`
    Links       []V1FileLink `json:"links"`
    Definitions *json.RawMessage `json:"definitions"`
}

func indentJson(input []byte) (output string, err error) {
  var parsedJson interface{}
  err = json.Unmarshal(input, &parsedJson)
  if err != nil {
    return
  }

  outputBytes,err := json.MarshalIndent(parsedJson, "", "  ")
  if err != nil {
    return
  }

  output = string(outputBytes)
  return
}

func ValidateResponse(schemaPath string, jsonBody interface{}) (bool, error) {
    if len(schemaPath) == 0 {
        err := fmt.Errorf("SchemaPath missing")
        return false, err
    }

    schemaFile,err := os.Open(schemaPath)
    if err != nil {
        fmt.Printf("schemaFile open error: %s\n", err)
        return false, err
    }

    var v1file V1File
    err = json.NewDecoder(schemaFile).Decode(&v1file)
    if err != nil{
        fmt.Printf("\nv1file err: %s\n", err)
    }

    resSchema := v1file.Links[0].ResponseSchema
    respSchemaString, err := json.Marshal(&resSchema)
    if err != nil {
        fmt.Printf("unmarshal err: %s\n", err)
    }

    indentedReqSchemaStr,err := indentJson(respSchemaString)
    if err != nil {
        fmt.Printf("indentedReqSchemaStr err: %s", err)
        return false, err
    }

    responseLoader := gojsonschema.NewGoLoader(jsonBody)
    schemaLoader := gojsonschema.NewStringLoader(indentedReqSchemaStr)

    result, err := gojsonschema.Validate(schemaLoader, responseLoader)
    if err != nil {
        fmt.Errorf("validate error: %s", err)
        return false, err
    }
    if result.Valid() {
        scamp.Info.Printf("The JSON response is valid\n")
    } else {
        fmt.Errorf("The document is not valid. see errors :\n")
        for _, err := range result.Errors() {
            fmt.Errorf("- %s\n", err)
        }
        return false, err
    }
    return true, err
}
