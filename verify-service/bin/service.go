package main

import (
  "os"
  "math"

  "bytes"
  // "flags"
  "encoding/json"
  "github.com/gudtech/scamp-go/scamp"
)

type VerifyRequest struct {
  TargetUrl string `json:"target_url"`
  SupportedActions []string `json:"supported_actions"`
}

type SqrtRequest struct {
  Input float64 `json:"input"`
}

type SqrtResponse struct {
  Output float64 `json:"output"`
  NaN bool `json:"nan"`
}

func main() {
  scamp.Initialize("/etc/GTSOA/soa.conf")

  svc,err := scamp.NewService("0.0.0.0:63531","service_verifier")
  if err != nil {
    scamp.Error.Printf("could not create service: `%s`", err.Error())
    os.Exit(1)
  }

  svc.Register("verify", func(msg *scamp.Message, client *scamp.Client) {
    var err error
    // scamp.Error.Printf("received new verify request size: %d", len(msg.Bytes()))

    var req VerifyRequest
    err = json.NewDecoder(bytes.NewReader(msg.Bytes())).Decode(&req)
    if err != nil {
      panic(err.Error())
    }

    // scamp.Info.Printf("%s", req)

    resp := scamp.NewResponseMessage()
    resp.SetRequestId(msg.RequestId)
    _,err = client.Send(resp)
    if err != nil {
      panic(err.Error())
    }
  })


  svc.Register("sqrt", func(msg *scamp.Message, client *scamp.Client) {
    var sqrtReq SqrtRequest
    err = json.NewDecoder(bytes.NewReader(msg.Bytes())).Decode(&sqrtReq)
    if err != nil {
      scamp.Error.Printf("nooo: `%s`", err)
      return
    }

    // scamp.Info.Printf("sqrt req: %s", sqrtReq)

    var sqrtResp SqrtResponse

    res := math.Sqrt(sqrtReq.Input)
    if math.IsNaN(res) {
      sqrtResp.NaN = true
    } else {
      sqrtResp.Output = res
    }

    var buf bytes.Buffer
    err = json.NewEncoder(&buf).Encode(sqrtResp)
    if err != nil {
      scamp.Error.Printf("nooo: `%s`", err)
      return
    }


    resp := scamp.NewResponseMessage()
    resp.SetRequestId(msg.RequestId)
    resp.Write(buf.Bytes())
    _,err = client.Send(resp)
    if err != nil {
      panic(err.Error())
    }
  })

  svc.Run()
}