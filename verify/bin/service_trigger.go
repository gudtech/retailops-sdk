package main

import (
  // "encoding/json"

  "github.com/gudTECH/scamp-go/scamp"
  // "bytes"

  // "os"
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
  msg.SetRequestId(1 /*reqId*/)
  scamp.Info.Printf("reqId: %d", 1/* reqId */)

  var req = map[string]string {}

  msg.WriteJson(req)
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

  scamp.Info.Printf("success")
}
