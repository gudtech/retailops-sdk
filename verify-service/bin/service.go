package main

import (
  "bytes"
  // "flags"
  "encoding/json"
  "github.com/gudtech/scamp-go/scamp"
)

type VerifyRequest struct {
  TargetUrl string `json:"target_url"`
  SupportedActions []string `json:"supported_actions"`
}

func main() {
  scamp.Initialize("/etc/GTSOA/soa.conf")

  svc,err := scamp.NewService("0.0.0.0:63530","service_verifier")
  if err != nil {
    scamp.Error.Printf("could not create service: `%s`", err.Error())
  }

  svc.Register("verify", func(msg *scamp.Message, client *scamp.Client) {
    var err error
    scamp.Error.Printf("received new verify request size: %d", len(msg.Bytes()))

    var req VerifyRequest
    err = json.NewDecoder(bytes.NewReader(msg.Bytes())).Decode(&req)
    if err != nil {
      panic(err.Error())
    }

    scamp.Info.Printf("%s", req)

    resp := scamp.NewResponseMessage()
    resp.SetRequestId(msg.RequestId)
    _,err = client.Send(resp)
    if err != nil {
      panic(err.Error())
    }
  })

  svc.Run()
}