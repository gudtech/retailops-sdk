package main

import (
  "os"
  "fmt"
  "flag"
  "github.com/gudtech/scamp-go/scamp"
  "sync"

  "github.com/gudtech/retailops-sdk/verify/verify_service"
  "github.com/gudtech/retailops-sdk/verify/sdk_actions"
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

  verifierSvc,err := scamp.NewService("main","0.0.0.0:0","sdk_service")
  if err != nil {
    scamp.Error.Printf("could not create service: `%s`", err.Error())
    os.Exit(1)
  }
  verify_service.TicketPath = *registrationTicket
  verifierSvc.Register("Integration.Channel.certify", verify_service.VerifyAction)

  callbackSvc,err := scamp.NewService("channelmodule", "0.0.0.0:0","sdk_service")
  if err != nil {
    scamp.Error.Printf("could not create serivce: `%s`", err.Error())
    os.Exit(1)
  }
  // callbackSvc.Register("SDK.items_returned",sdk_actions.OrderReturnedV1) // , SDK callback: order_returned_v1
  // callbackSvc.Register("SDK.shipment_submit",sdk_actions.OrderShipmentSubmitV1) // , SDK callback: order_shipment_submit_v1
  // callbackSvc.Register("SDK.order_cancel",sdk_actions.OrderCancelV1) // , SDK callback: order_cancel_v1
  // callbackSvc.Register("SDK.order_complete",sdk_actions.OrderCompleteV1) // , SDK callback: order_complete_v1
  // callbackSvc.Register("SDK.capture_channel_payments",sdk_actions.OrderSettlePaymentV1) // , SDK callback: order_settle_payment_v1
  // callbackSvc.Register("SDK.orderpull.order_ack",sdk_actions.OrderAcknowledgeV1) // , SDK callback: order_acknowledge_v1
  callbackSvc.Register("SDK.orderpull.order_fetch",sdk_actions.OrderPullV1) // , SDK callback: order_pull_v1
  // callbackSvc.Register("SDK.catpush_config",sdk_actions.CatalogGetConfigV1) // , SDK callback: catalog_get_config_v1
  // callbackSvc.Register("SDK.catpush_transmit",sdk_actions.CatalogPushV1) // , SDK callback: catalog_push_v1
  callbackSvc.Register("SDK.inventory.invpush_transmit",sdk_actions.InventoryPushV1) // , SDK callback: inventory_push_v1
  // callbackSvc.Register("SDK.writeback",sdk_actions.OrderUpdateV1) // , SDK callback: order_update_v1

  announcer,err := scamp.NewDiscoveryAnnouncer()
  if err != nil {
    scamp.Error.Printf("failed to create announcer: `%s`", err)
    return
  }

  announcer.Track(verifierSvc)
  announcer.Track(callbackSvc)
  go announcer.AnnounceLoop()

  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
    callbackSvc.Run()
    wg.Done()
  }()

  wg.Add(1)
  go func(){
    verifierSvc.Run()
    wg.Done()
  }()

  wg.Wait()
}
