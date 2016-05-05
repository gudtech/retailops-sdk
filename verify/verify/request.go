package verify

import (
  "fmt"
  "io"
  "io/ioutil"

  "strings"

  "encoding/json"

  "net/http"
  "net/url"

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

func Request(baseUrlStr string, hyperSchema io.Reader, example io.Reader, verbose bool) (err error) {
  requestUrl,err := url.Parse(baseUrlStr)
  if err != nil {
    return
  }

  basePath := requestUrl.Path

  var v1file V1File
  err = json.NewDecoder(hyperSchema).Decode(&v1file)

  exampleBytes,err := ioutil.ReadAll(example)
  if err != nil {
    return
  }

  for _,link := range v1file.Links {
    err = requestAgainstLink(v1file, link, basePath, requestUrl, exampleBytes, verbose)
    if err != nil {
      return err
    }
  }
  return
}

func requestAgainstLink(v1file V1File, link V1FileLink, basePath string, requestUrl *url.URL, exampleBytes []byte, verbose bool) (err error) {
  client := &http.Client{}

  /*
    schema lib data setup
  */
  indentedReqSchemaStr,err := indentJson(*link.RequestSchema)
  if err != nil {
    return err
  }

  reqSchemaWDefs,err := insertDefinitions(v1file.Definitions, indentedReqSchemaStr)
  if err != nil {
    return err
  }
  reqSchemaLoader := schema.NewStringLoader(reqSchemaWDefs)

  respSchemaStr := string(*link.ResponseSchema)
  respSchemaWDefs,err := insertDefinitions(v1file.Definitions, respSchemaStr)
  if err != nil {
    return err
  }
  respSchemaLoader := schema.NewStringLoader(respSchemaWDefs)

  exampleStr := string(exampleBytes)
  exampleDataLoader := schema.NewStringLoader(exampleStr)

  /*
    Echo request to be performed
  */
  if verbose {
    fmt.Println("HTTP request:", strings.ToUpper(link.Method), fmt.Sprintf("%s%s", basePath, link.Href))

    // fmt.Println("schema (with definitions omitted):")
    // fmt.Println(indentedReqSchemaStr)
    fmt.Println("HTTP request body:")
    fmt.Println(exampleStr)
  }

  result,err := schema.Validate(reqSchemaLoader, exampleDataLoader)
  if err != nil {
    fmt.Println("error validating:", err.Error())
    return err
  } else if !result.Valid() {
    // TODO: iterate over result errors
    fmt.Printf("validation failures:\n")
    for _,resultErr := range result.Errors() {
      fmt.Printf("- %s\n", resultErr)
    }
    fmt.Println()
    return fmt.Errorf("outgoing example invalid")
  }

  // fmt.Println("example is valid:", result.Valid())

  /* 
    HTTP request issuance
  */
  requestUrl.Path = fmt.Sprintf("%s%s", basePath, link.Href)
  response,err := client.Do(&http.Request {
    Method: strings.ToUpper(link.Method),
    URL: requestUrl,
  })
  if err != nil {
    return err
  }


  /*
    HTTP response validation
  */
  responseBytes,err := ioutil.ReadAll(response.Body)
  if response.StatusCode != 200 {
    err = fmt.Errorf("HTTP status code: %d", response.StatusCode)
    return err
  }

  // fmt.Println("response status code:", response.StatusCode)
  indentedResp,err := indentJson(responseBytes)
  if err != nil {
    return err
  }
  if verbose {
    fmt.Println("HTTP response status code:",response.StatusCode)
    fmt.Println("HTTP response body:")
    fmt.Println(indentedResp)
  }

  respDataLoader := schema.NewStringLoader(string(responseBytes))
  result,err = schema.Validate(respSchemaLoader, respDataLoader)
  if verbose {
    fmt.Println("response valid:", result.Valid())
  }
  if !result.Valid() {
    fmt.Println("reason(s):")
    for _,validationError := range result.Errors() {
      fmt.Println(" ",validationError)
    }
    return fmt.Errorf("invalid response")
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