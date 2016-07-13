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

  verifierSvc,err := scamp.NewService("main","0.0.0.0:","sdk_service")
  if err != nil {
    scamp.Error.Printf("could not create service: `%s`", err.Error())
    os.Exit(1)
  }
  verify_service.TicketPath = *registrationTicket
  verifierSvc.Register("Integration.Channel.certify", verify_service.VerifyAction)

  callbackSvc,err := scamp.NewService("sdk", "127.0.0.1:6000","sdk_service")
  if err != nil {
    scamp.Error.Printf("could not create serivce: `%s`", err.Error())
    os.Exit(1)
  }
  callbackSvc.Register("SDK.items_returned",sdk_actions.ItemsReturnedV1) // , SDK callback: order_returned
  callbackSvc.Register("SDK.shipment_submit",sdk_actions.ShipmentSubmitV1) // , SDK callback: order_shipment_submit
  callbackSvc.Register("SDK.writeback",sdk_actions.WritebackV1) // , SDK callback: unknown
  callbackSvc.Register("SDK.order_cancel",sdk_actions.OrderCancelV1) // , SDK callback: order_cancel
  callbackSvc.Register("SDK.order_complete",sdk_actions.OrderCompleteV1) // , SDK callback: order_complete
  callbackSvc.Register("SDK.capture_channel_payments",sdk_actions.CaptureChannelPaymentsV1) // , SDK callback:
  callbackSvc.Register("SDK.order_ack",sdk_actions.OrderAckV1) // , SDK callback: order_acknowledge
  callbackSvc.Register("SDK.order_fetch",sdk_actions.OrderFetchV1) // , SDK callback: order_pull
  callbackSvc.Register("SDK.catpush_config",sdk_actions.CatpushConfigV1) // , SDK callback: catalog_get_config
  callbackSvc.Register("SDK.catpush_transmit",sdk_actions.CatpushTransmitV1) // , SDK callback: catalog_push
  callbackSvc.Register("SDK.invpush_transmit",sdk_actions.InvpushTransmitV1) // , SDK callback: inventory_push
  callbackSvc.Register("SDK.order_update",sdk_actions.OrderUpdateV1) // , SDK callback: order_update
  callbackSvc.Register("SDK.order_settle_payment",sdk_actions.OrderSettlePaymentV1) // , SDK callback: order_settle_payment

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
