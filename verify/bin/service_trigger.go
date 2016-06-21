package main

import (
  "encoding/json"

  "github.com/gudtech/scamp-go/scamp"
  "github.com/gudtech/retailops-sdk/verify/common"
  "bytes"

  "os"
)

func main() {
  scamp.Initialize("/etc/GTSOA/soa.conf")

  client,err := scamp.Dial("0.0.0.0:63531")
  msg := scamp.NewRequestMessage()
  msg.SetAction("verify")
  msg.SetRequestId(1 /*reqId*/)
  scamp.Info.Printf("reqId: %d", 1/* reqId */)

  var req = common.VerifyRequest {
    Version: 1,
    TargetUrl: "http://localhost:5000/api/channel",
    SupportedActions: []string{ "catalog_get_config" },
  }

  msg.WriteJson(req)
  respChan,err := client.Send(msg)
  if err != nil {
    scamp.Error.Printf("could not send message: `%s`", err)
  }
  respMsg := <-respChan

  var resp common.VerifyResponse
  err = json.NewDecoder(bytes.NewReader(respMsg.Bytes())).Decode(&resp)
  if err != nil {
    scamp.Error.Printf("could not decode response: %s", err.Error())
    os.Exit(1)
  }

  if resp.Status != "success" {
    scamp.Error.Printf("failed!")
    os.Exit(1)
  }

  scamp.Info.Printf("success")
}