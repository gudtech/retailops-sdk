package main

import (
  "os"
  "fmt"
  // "flag"
  "github.com/gudtech/scamp-go/scamp"

  "github.com/gudtech/retailops-sdk/verify-service/verify_service"
)



func main() {
  var err error

  err = scamp.Initialize("/etc/GTSOA/soa.conf")
  if err != nil {
    fmt.Println("failed to load scamp config:",err.Error())
    os.Exit(1)
  }

  verifierSvc,err := scamp.NewService("0.0.0.0:63531","service_verifier")
  if err != nil {
    scamp.Error.Printf("could not create service: `%s`", err.Error())
    os.Exit(1)
  }

  /*
  * {
  *   "target_url": "http://bob.com/ro_api",
  *   "supported_actions": ["asdf", "fdsa"]
  * }
  */
  verifierSvc.Register("verify", verify_service.VerifyAction)

  verifierSvc.Run()
}