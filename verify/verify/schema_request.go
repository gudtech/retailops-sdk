package verify

import (
  "fmt"
  "io"
  "io/ioutil"

  "strings"
  "bytes"

  "encoding/json"

  "net/http"
  "net/url"

  "time"

  schema "github.com/xeipuuv/gojsonschema"
)

var client = &http.Client{
  Timeout: time.Second * 15,
}

/*
type V1FileAction struct {
  Description string   `json:"description"`
  Example     string   `json:"example"`
  Type        []string `json:"type"`
}

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

func Request(baseUrlStr, integrationAuthKey string, hyperSchema io.Reader, example io.Reader, verbose bool, expectedStatusResponseCode int) (err error) {

  requestUrl,err := url.Parse(baseUrlStr)
  if err != nil {
    return
  }

  basePath := requestUrl.Path

  var v1file V1File
  // fmt.Errorf("v1file: ", v1file)
  err = json.NewDecoder(hyperSchema).Decode(&v1file)
  if err != nil{
      fmt.Println("\nv1file err:", err)
  }

  exampleBytes,err := ioutil.ReadAll(example)
  if err != nil {
      fmt.Errorf("error: %s ", err)
    return
  }
  // fmt.Println("exampleBytes", exampleBytes)
  exampleBytes,err = insertIntegrationAuthKey(exampleBytes, integrationAuthKey)
  if err != nil {
    return
  }

  // fmt.Println("exampleBytes 2", exampleBytes)

  for _,link := range v1file.Links {
    err = requestAgainstLink(v1file, link, basePath, requestUrl, exampleBytes, verbose, expectedStatusResponseCode)
    if err != nil {
      return err
    }
  }
  return
}

func requestAgainstLink(v1file V1File, link V1FileLink, basePath string, requestUrl *url.URL, exampleBytes []byte, verbose bool, expectedStatusResponseCode int) (err error) {
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
    indentedExampleSchemaStr,err := indentJson(exampleBytes)
    if err != nil {
      return err
    }

    fmt.Println(strings.ToUpper(link.Method), fmt.Sprintf("%s%s", basePath, link.Href))
    // fmt.Println("schema (with definitions omitted):")
    // fmt.Println(indentedReqSchemaStr)
    fmt.Println("HTTP request body:")
    fmt.Println(indentedExampleSchemaStr)
  }

  result,err := schema.Validate(reqSchemaLoader, exampleDataLoader)
  if err != nil {
    fmt.Println("error validating:", err.Error())
    return err
  } else if !result.Valid() {
      // TODO: iterate over result errors
        var buf bytes.Buffer
        _,err := buf.WriteString("\n\nFailure validating outgoing test:\n")
        if err != nil {
          return err
        }

        for _,resultErr := range result.Errors() {
          _,err := buf.WriteString(fmt.Sprintf("- %s\n", resultErr))
          if err != nil {
            return err
          }
        }

    return fmt.Errorf(buf.String())
  }

  // fmt.Println("example is valid:", result.Valid())

  /*
    HTTP request issuance
  */
  requestBuf := bytes.NewBuffer(exampleBytes)
  requestUrl.Path = fmt.Sprintf("%s%s", basePath, link.Href)

  request,err := http.NewRequest(
    strings.ToUpper(link.Method),
    requestUrl.String(),
    requestBuf,
  )
  if err != nil {
    return
  }
  //set Content-Type
  request.Header.Set("Content-Type", "application/json")

  response,err := client.Do(request)
  if err != nil {
    return err
  }


  /*
    HTTP response validation
  */
  responseBytes,err := ioutil.ReadAll(response.Body)
  // if response.StatusCode != 200 {
  //expected responses should be 401 when we test by sending a bad key
  if response.StatusCode != expectedStatusResponseCode {
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
    var buf bytes.Buffer
    _,err := buf.WriteString("reason(s):\n")
    if err != nil {
      return err
    }

    for _,validationError := range result.Errors() {
      _,err := buf.WriteString(fmt.Sprintf(" %s",validationError))
      if err != nil {
        return err
      }
    }
    return fmt.Errorf(buf.String())
  }

  return
}

func insertIntegrationAuthKey(in []byte, integrationAuthKey string) (out []byte, err error) {
  var asMap map[string]interface{}
  inReader := bytes.NewReader(in)
  err = json.NewDecoder(inReader).Decode(&asMap)
  if err != nil {
    return
  }

  asMap["integration_auth_token"] = integrationAuthKey

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(asMap)
  if err != nil {
    return
  }

  out = buf.Bytes()

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
