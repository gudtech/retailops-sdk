package faker

import (
  "fmt"
  "io"
  "io/ioutil"

  "encoding/json"
)

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

/* 
  Outgoing request
   * Links
    * HREF/Method/Schema of outgoing data
    * TargetSchema response format (pull payload from ex file)
*/
type V1File struct {
  Schema     string `json:"$schema"`
  Links      []V1FileLink `json:"links"`
  Definitions V1FileDefinitions `json:"definitions"`
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
    fmt.Println(link.Method, link.Href, string(*link.RequestSchema), string(examplesBytes))
  }
  // fmt.Println(v1file.Links)
  return
}