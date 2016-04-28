package verify

import (
  "fmt"
  "io"
  "io/ioutil"

  // "strings"

  "encoding/json"

  schema "github.com/xeipuuv/gojsonschema"
)

type V1FileAction struct {
  Description string   `json:"description"`
  Example     string   `json:"example"`
  Type        []string `json:"type"`
}

/*
type V1FileDefinitions struct {
  Action V1FileAction `json:"action"`
  ResponseSchema *json.RawMessage `json:"event"`
  Identity struct{} `json:"identity"`
  Version  struct {
    Description string   `json:"description"`
    Example     int      `json:"example"`
    Type        []string `json:"type"`
  } `json:"version"`
}
*/

/* 
  Outgoing request
   * Links
    * HREF/Method/Schema of outgoing data
    * TargetSchema response format (pull payload from ex file)
*/
type V1File struct {
  Schema      string `json:"$schema"`
  Links       []V1FileLink `json:"links"`
  Definitions *json.RawMessage `json:"definitions"`
  // Properties struct{} `json:"properties"`
  // Stability  string   `json:"stability"`
  // Title      string   `json:"title"`
  // Type       []string `json:"type"`
}

type V1FileLink struct {
  Description    string `json:"description"`
  Href           string `json:"href"`
  Method         string `json:"method"`
  Rel            string `json:"rel"`
  RequestSchema  *json.RawMessage `json:"schema"`
  ResponseSchema *json.RawMessage `json:"targetSchema"`
  Title          string `json:"title"`
}

func Request(hyperSchema io.Reader, example io.Reader) (err error) {
  var v1file V1File
  err = json.NewDecoder(hyperSchema).Decode(&v1file)

  examplesBytes,err := ioutil.ReadAll(example)
  if err != nil {
    return
  }

  fmt.Println(len(v1file.Links), "REQUEST(S) TO BE GENERATED", )
  for _,link := range v1file.Links {
    reqSchemaStr := string(*link.RequestSchema)
    schemaWithDefinitions,err := insertDefinitions(v1file.Definitions, reqSchemaStr)
    if err != nil {
      return err
    }


    exampleStr := string(examplesBytes)

    reqSchemaLoader := schema.NewStringLoader(schemaWithDefinitions)
    exampleDataLoader := schema.NewStringLoader(exampleStr)

    fmt.Println(link.Method, link.Href)
    fmt.Println("schema (with definitions omitted):")
    fmt.Println(reqSchemaStr)
    fmt.Println("example:")
    fmt.Println(exampleStr)

    result,err := schema.Validate(reqSchemaLoader, exampleDataLoader)
    if err != nil {
      fmt.Println("error validating:", err.Error())
      return err
    } else if !result.Valid() {
      err = fmt.Errorf("example was not valid: ")
    }

    fmt.Println("example is valid:", result.Valid())
  }
  return
}

func insertDefinitions(rawDefinitions *json.RawMessage, schemaStr string) (fixedupSchema string, err error) {
  var schemaDoc map[string]interface{}
  err = json.Unmarshal([]byte(schemaStr), &schemaDoc)
  if err != nil {
    return
  }

  schemaDoc["definitions"] = rawDefinitions
  fixedupSchemaBytes,err := json.MarshalIndent(schemaDoc, "", "  ")
  if err != nil {
    return
  }

  fixedupSchema = string(fixedupSchemaBytes)
  return
}