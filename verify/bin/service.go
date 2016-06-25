package main

import (
  "os"
  "fmt"
  "flag"
  "github.com/gudtech/scamp-go/scamp"

  "github.com/gudtech/retailops-sdk/verify/verify_service"
)

var gtsoaConfigPath = flag.String("config", "/etc/GTSOA/soa.conf", "path to the GTSOA soa.conf")
var registrationTicket = flag.String("registration-ticket", "/backplane/etc/endorsements/gtuser-sdk_service.ticket", "path to gtuser sdk_service ticket")

func main() {
  var err error
  flag.Parse()

  fmt.Println("loading config from",*gtsoaConfigPath)

  err = scamp.Initialize(*gtsoaConfigPath)
  if err != nil {
    fmt.Println("failed to load scamp config:",err.Error())
    os.Exit(1)
  }

  verifierSvc,err := scamp.NewService("main","0.0.0.0:","sdk_service")
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

  verify_service.TicketPath = *registrationTicket
  verifierSvc.Register("Integration.Channel.certify", verify_service.VerifyAction)

  announcer,err := scamp.NewDiscoveryAnnouncer()
  if err != nil {
    scamp.Error.Printf("failed to create announcer: `%s`", err)
    return
  }
  announcer.Track(verifierSvc)
  go announcer.AnnounceLoop()

  verifierSvc.Run()
}