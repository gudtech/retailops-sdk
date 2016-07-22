package main

import (
  // "encoding/json"
  "github.com/gudTECH/scamp-go/scamp"
)

func main() {
  scamp.Initialize("/etc/GTSOA/soa.conf")
  scamp.Info.Printf("dialing")
  client,err := scamp.Dial("127.0.0.1:6000")
  if err != nil {
      scamp.Info.Printf("err: ", err)
  }
  msg := scamp.NewRequestMessage()
  msg.SetAction("SDK.invpush_transmit")
  // msg.SetAction("SDK.capture_channel_payments")
  // msg.SetAction("SDK.writeback")

  msg.SetRequestId(1 /*reqId*/)
  scamp.Info.Printf("reqId: %d", 1/* reqId */)

  var req = []byte (`{
    "headers": {
      "client_id": 1,
      "ticket": "RETAILOPS_SDK"
    },
    "version": 1,
    "action": "Mozu.inventory.invpush_transmit",
    "data": {
      "inventory": {
        "data": [
          {
            "sku": "53",
            "id": "66",
            "concept": "sku",
            "qty_breakdown": [
              {
                "sku": "53",
                "sellable": 5,
                "unclaimed": 0,
                "zones": [
                  {
                    "zone": "Default",
                    "pick": 0,
                    "npick": 5
                  }
                ],
                "facility": "Testoria"
              },
              {
                "est_ship": "2016-04-22T11:02:47-07:00",
                "sku": "53",
                "reserving_orders": [
                  "2390"
                ],
                "sellable": 32,
                "unclaimed": 24,
                "vendor": "002_Acme Corp"
              }
            ],
            "qty_available": 24
          }
        ]
      },
      "client_id": 1,
      "channel": {
        "params": {
          "breakdown_inventory": 0,
          "tenant": "15394",
          "appKey": "dripclub.retailops.0.1.0.Release",
          "base_uri": "http://localhost:5000/api/channel"
        },
        "id": 30
      }
    }
  }
`)

  msg.Write(req)
  respChan,err := client.Send(msg)
  if err != nil {
    scamp.Error.Printf("could not send message: `%s`", err)
  }
  scamp.Info.Printf("waiting")
  respMsg := <-respChan
  scamp.Info.Printf("response: %s", string(respMsg.Bytes()))
  // var resp common.VerifyResponse
  // err = json.NewDecoder(bytes.NewReader(respMsg.Bytes())).Decode(&resp)
  // if err != nil {
  //   scamp.Error.Printf("could not decode response: %s", err.Error())
  //   os.Exit(1)
  // }
  //
  // if resp.Status != "success" {
  //   scamp.Error.Printf("failed!")
  //   os.Exit(1)
  // }

  scamp.Info.Printf("response received")
  // }
}
