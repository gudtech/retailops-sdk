package main

import (
  "flag"
  "fmt"
  "strings"

  // "net/http"
  // "encoding/json"
  "os"
  // p "path"
  // fp "path/filepath"

  "github.com/gudtech/retailops-sdk/verify/verify"
)

type schemaExample struct {
  schemaPath string
  examplePath string
}

var HR string = "----------------"

func main() {
  var cliExec verify.CLIExecution
  var err error

  schemaPathPtr := flag.String("schema-path", "", "path to JSON or directory with JSON")
  baseURLPtr := flag.String("base-url", "http://localhost:5000/api/channel", "base url for sending requests")
  stopOnError := flag.Bool("stop-on-error", false, "stop immediately on error")
  filterPtr := flag.String("filter", "", "filter test cases by name. ex: 'order' or 'order_cance'")
  verbosePtr := flag.Bool("verbose", false, "show all outgoing and incoming request data")
  apiKeyPtr := flag.String("api-key", "", "your retailops API key")
  certifyActionsPtr := flag.String("certify-actions", "catalog_get_config,catalog_push,inventory_push,order_acknowledge,order_cancel,order_complete,order_pull,order_returned,order_settle_payment,order_shipment_submit,order_update", "subset of actions to test for certification")

  flag.Parse()

  args := flag.Args()
  if len(args) > 0 && (args[0] == "certify") {
    cliExec.Action = "certify"
  } else if len(args) > 0 {
    fmt.Println("unknown action", args[0])
    os.Exit(1)
  } else {
    cliExec.Action = "test"
  }

  if len(*baseURLPtr) == 0 {
    fmt.Println("base-url cannot be empty")
    os.Exit(1)
  }
  cliExec.BaseURL = *baseURLPtr
  cliExec.SchemaPath = *schemaPathPtr
  cliExec.SchemaPathIsDir = isDir(*schemaPathPtr)
  cliExec.StopOnError = *stopOnError
  cliExec.SchemaFilter = *filterPtr
  cliExec.Verbose = *verbosePtr
  cliExec.ApiKey = *apiKeyPtr

  if len(*certifyActionsPtr) != 0 {
    cliExec.CertifyActions = strings.SplitN(*certifyActionsPtr,",",-1)
  }

  err = verify.Execute(cliExec)
  if err != nil {
    fmt.Println("failed:", err.Error())
    os.Exit(1)
  } else {
    if cliExec.Action == "test" {
      fmt.Println("local verify was successful")
    } else if cliExec.Action == "certify" {
      fmt.Println("remote certification was a success")
    }
  }
}

func isDir(path string) (bool) {
  info, err := os.Stat(path)
  return err == nil && info.IsDir()
}