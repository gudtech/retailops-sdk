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
  // msg.SetAction("SDK.invpush_transmit")
  // msg.SetAction("SDK.capture_channel_payments")
  // msg.SetAction("SDK.writeback")
  msg.SetAction("SDK.order_ack")

  msg.SetRequestId(1 /*reqId*/)
  scamp.Info.Printf("reqId: %d", 1/* reqId */)

  var req = []byte (`{
    "headers": {
      "client_id": 1,
      "ticket": "RETAILOPS_SDK"
    },
    "version": 1,
    "action": "Aabaco.orderpull.order_ack",
    "data": {
      "order": {
        "acks": [
          "496"
        ]
      },
      "client_id": 1,
      "channel": {
        "params": {
          "StoreID": "yhst-18909142938879050075142",
          "next_order_refnum": 496,
          "order_ack_status_id": "32",
          "order_fulfilled_status_id": "34",
          "order_in_filfillment_status_id": "33",
          "base_uri": "http://localhost:5000/api/channel"
        },
        "id": 21
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

  scamp.Info.Printf("success: response received")
  // }
}
