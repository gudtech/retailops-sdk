package main

import (
  "time"
  "github.com/gudtech/scamp-go/scamp"
)

type VerifyRequest struct {
  TargetUrl string `json:"target_url"`
  SupportedActions []string `json:"supported_actions"`
}

func main() {
  scamp.Initialize("/etc/GTSOA/soa.conf")

  client,err := scamp.Dial("0.0.0.0:63530")
  if err != nil {
    scamp.Error.Printf("could not dial service: `%s`", err)
  }

  for i := 0; i < 8000; i++ {
    reqId := i+1
    go func() {
      msg := scamp.NewRequestMessage()
      msg.SetAction("verify")
      scamp.Info.Printf("reqId: %d", reqId)
      msg.SetRequestId(reqId)
      msg.WriteJson(VerifyRequest {
        TargetUrl: "http://localhost:9200",
        SupportedActions: []string { "asdf", "fdas" },
      })
      respChan,err := client.Send(msg)
      if err != nil {
        scamp.Error.Printf("could not send message: `%s`", err)
      }
      <-respChan
    }()
  }

  time.Sleep(time.Duration(1) * time.Hour)

}