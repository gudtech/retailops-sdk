package verify

import (
  "os"

  "fmt"
  "bytes"
  "encoding/json"

  p "path"
  fp "path/filepath"
  "strings"

  "net/http"
  "time"

  "github.com/gudtech/retailops-sdk/verify/common"
)

type SchemaExample struct {
  SchemaPath string
  ExamplePath string
}

type CLIExecution struct {
  Action string
  SchemaPath string
  BaseURL string
  SchemaFilter string

  ApiKey string

  SchemaPathIsDir bool
  StopOnError bool
  Verbose bool
}

var HR string = "----------------"

func Execute(cliExec CLIExecution) (err error) {
  if cliExec.Action == "test" {
    return doLocalTest(cliExec)
  } else if cliExec.Action == "certify" {
    return doCertify(cliExec)
  }

  return
}

var certifyClient = &http.Client{
  Timeout: time.Minute * 5,
}

var baseDispatcherUrl string = "https://api.retailops.com/integrations/channel/certify.json"
func doCertify(cliExec CLIExecution) (err error) {
  var verReq = common.VerifyRequest {
    Version: 1,
    TargetUrl: cliExec.BaseURL,
    SupportedActions: []string{ "catalog_get_config" },
  }

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(verReq)
  if err != nil {
    return
  }

  url := fmt.Sprintf("%s?apikey=%s",baseDispatcherUrl,cliExec.ApiKey)
  resp,err := certifyClient.Post(url, "application/json", &buf)
  if err != nil {
    return
  }


  var apiResp common.VerifyResponse
  err = json.NewDecoder(resp.Body).Decode(&apiResp)
  if err != nil {
    return
  }

  if apiResp.Status != "success" {
    err = fmt.Errorf("certification was not successful: %s", apiResp.Message)
  }

  return

}

func doLocalTest(cliExec CLIExecution) (err error) {
  var examples []SchemaExample
  if cliExec.SchemaPathIsDir {
    examples,err = allExamples(cliExec.SchemaPath, cliExec.SchemaFilter)
  } else {
    examples,err = examplesForSchema(cliExec.SchemaPath)
  }

  if err != nil {
    return
  } else if len(examples) == 0 {
    err = fmt.Errorf("no verification files found. try `--help` for more information")
    return
  }

  fmt.Println(len(examples),"TESTS TO BE GENERATED")
  for index,example := range examples {
    doVerify(cliExec, index, example)
  }

  return
}


func doVerify(cliExec CLIExecution, index int, testCase SchemaExample) (err error) {
  var thereWasAnError bool

  fmt.Println(HR)
  fmt.Printf("TEST %d (%s)", index+1, p.Base(testCase.ExamplePath))
  err = request(cliExec.BaseURL, testCase.SchemaPath, testCase.ExamplePath, cliExec.Verbose)
  if err != nil {
    fmt.Printf("\rTEST %d (%s) FAILED: %s\n", index+1, p.Base(testCase.ExamplePath), err.Error())
    if cliExec.StopOnError {
      os.Exit(1)
    } else {
      thereWasAnError = true
    }
  } else {
    if cliExec.Verbose {
      fmt.Println("")
    }
    fmt.Printf("\rTEST %d (%s) WAS A SUCCESS\n", index+1, p.Base(testCase.ExamplePath))
    if cliExec.Verbose {
      fmt.Println("")
    }
  }

  if thereWasAnError && cliExec.StopOnError {
    err = fmt.Errorf("AT LEAST ONE OF THE TEST CASES FAILED")
    return
  }

  return
}

func request(baseUrl, schemaPath, examplePath string, verbose bool) (err error) {
  f,err := os.Open(schemaPath)
  if err != nil {
    return
  }

  exampleF,err := os.Open(examplePath)
  if err != nil {
    return
  }

  err = Request(baseUrl, f, exampleF, verbose)
  if err != nil {
    return
  }

  return
}

func allExamples(dirname, filter string) (verifications []SchemaExample, err error) {
  verifications = make([]SchemaExample,0)
  var allSchemasGlob string
  if filter == "" {
    allSchemasGlob = p.Join(dirname, "*v1.json")
  } else {
    allSchemasGlob = p.Join(dirname, fmt.Sprintf("*%s*v1.json", filter))
  }
  allSchemaPaths,err := fp.Glob(allSchemasGlob)
  if err != nil {
    return
  }
  if len(allSchemaPaths) == 0 {
    err = fmt.Errorf("`%s` did not contain schemas", dirname)
    return
  }

  for _,schemaPath := range allSchemaPaths {
    // fmt.Println(schemaPath)
    exs,err := examplesForSchema(schemaPath)
    if err != nil {
      return nil,err
    }

    for _,ex := range exs {
      verifications = append(verifications, ex)
    }
  }

  return
}

func examplesForSchema(schemaPath string) (verifications []SchemaExample, err error) {
  dirname,filename := p.Split(schemaPath)
  exampleFilename := strings.Replace(filename, ".json", "", -1)

  exampleFilenameGlob := fmt.Sprintf("%s_ex_*.json", exampleFilename)
  pathGlob := p.Join(dirname, exampleFilenameGlob)

  examplePaths,err := fp.Glob(pathGlob)
  if err != nil {
    return
  }

  verifications = make([]SchemaExample,0)
  for _,exPath := range examplePaths {
    verifications = append(verifications, SchemaExample{
      SchemaPath: schemaPath,
      ExamplePath: exPath,
    })
  }

  return
}