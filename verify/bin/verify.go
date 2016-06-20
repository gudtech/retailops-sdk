package main

import (
  "flag"
  "fmt"
  // "strings"

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

  flag.Parse()

  args := flag.Args()
  if len(args) != 1 {
    cliExec.Action = "test"
  } else if len(args) > 0 && args[0] == "certify" {
    cliExec.Action = "certify"
  } else {
    fmt.Println("unknown action", args[0])
    os.Exit(1)
  }

  if len(*schemaPathPtr) == 0 {
    fmt.Println("must set -schema-path")
    os.Exit(1)
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

  err = verify.Execute(cliExec)
  if err != nil {
    fmt.Println("failed:", err.Error())
    os.Exit(1)
  } else {
    fmt.Println("verify was successful")
  }
}

func isDir(path string) (bool) {
  info, err := os.Stat(path)
  return err == nil && info.IsDir()
}