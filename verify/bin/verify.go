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
  var cliExec = verify.CLIExecution{}
  var err error

  schemaPathPtr := flag.String("schema-path", "", "path to JSON or directory with JSON")
  baseURLPtr := flag.String("base-url", "http://localhost:5000/api/channel", "base url for sending requests")
  stopOnError := flag.Bool("stop-on-error", false, "stop immediately on error")
  filterPtr := flag.String("filter", "", "filter test cases by name. ex: 'order' or 'order_cancel'")
  verbosePtr := flag.Bool("verbose", false, "show all outgoing and incoming request data")
  apiKeyPtr := flag.String("api-key", "", "your retailops API key")
  certifyActionsPtr := flag.String("certify-actions", "catalog_get_config,catalog_push,inventory_push,order_acknowledge,order_cancel,order_complete,order_pull,order_returned,order_settle_payment,order_shipment_submit,order_update", "subset of actions to test for certification")
  integrationNamePtr := flag.String("integration-name", "", "human readable name for identifying the integration")
  roCertifyURLPtr := flag.String("retailops-certify-url", "https://api.retailops.com/integration/channel/certify.json", "")

  flag.Parse()

  args := flag.Args()
  if len(args) > 0 {
    action := args[0]
    if action == "certify" {
      cliExec.Action = "certify"
    } else if action == "show-auth-key" {
      cliExec.Action = "show_token"
    } else if action == "generate-auth-key" {
      cliExec.Action = "generate_token"
    } else if action == "install-auth-key" {
      if len(args) < 2 {
        fmt.Println("usage: verify install-token INTEGRATION_AUTH_KEY")
        os.Exit(1)
      }
      cliExec.Action = "install_token"
      cliExec.IntegrationAuthKey = args[1]
    // } else if action == "delete-token" {
    //   cliExec.Action = "delete_token"
    } else {
      fmt.Println("unknown action", action,"check usage for more information")
      os.Exit(1)
    }
  } else {
    cliExec.Action = "test"
  }

  if len(*baseURLPtr) == 0 {
    fmt.Println("base-url cannot be empty")
    os.Exit(1)
  }

  ats,err := verify.NewAuthTokenStorage()
  if err == nil {
    token,err := ats.ReadToken()
    if err == nil {
      cliExec.IntegrationAuthKey = token
    }
  }

  cliExec.BaseURL = *baseURLPtr
  cliExec.SchemaPath = *schemaPathPtr
  cliExec.SchemaPathIsDir = isDir(*schemaPathPtr)
  cliExec.StopOnError = *stopOnError
  cliExec.SchemaFilter = *filterPtr
  cliExec.Verbose = *verbosePtr
  cliExec.ApiKey = *apiKeyPtr
  cliExec.ROCertifyURL = *roCertifyURLPtr
  cliExec.IntegrationName = *integrationNamePtr

  if cliExec.Action == "install-token" {
    if len(cliExec.IntegrationAuthKey) < 20 {
      fmt.Println("integration auth key must be at least 20 characters")
      os.Exit(1)
    }
  }

  if len(*certifyActionsPtr) != 0 {
    cliExec.CertifyActions = strings.SplitN(*certifyActionsPtr,",",-1)
  }

  err = verify.Execute(cliExec)
  if err != nil {
    fmt.Println(err.Error())
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