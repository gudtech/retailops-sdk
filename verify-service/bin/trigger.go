package main

import (
  "sync"
  "github.com/gudtech/scamp-go/scamp"
  "math/rand"
  "math"

  "encoding/json"
  "bytes"
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

  var wg sync.WaitGroup

  for j := 0; j < 2; j++ {
    wg.Add(1)
    go func(){
      client,err := scamp.Dial("0.0.0.0:8080")
      if err != nil {
        scamp.Error.Printf("could not dial service: `%s`", err)
        return
      }
          scamp.Info.Printf("sup")

      for i := 0; i < 1; i++ {
        reqId := i+1
        wg.Add(1)
        go func() {
          msg := scamp.NewRequestMessage()
          msg.SetAction("sqrt")
          scamp.Info.Printf("reqId: %d", reqId)

          var req SqrtRequest
          req.Input = rand.NormFloat64()

          expSqrt := math.Sqrt(req.Input)

          var expResp SqrtResponse
          if math.IsNaN(expSqrt) {
            expResp.NaN = true
          } else {
            expResp.Output = expSqrt
          }
          expBytes,err := json.Marshal(expResp)
          if err != nil {
            scamp.Info.Printf("could not marshal exp resp: `%s`", err)
            return
          }

          msg.SetRequestId(reqId)
          msg.WriteJson(req)
          respChan,err := client.Send(msg)
          if err != nil {
            scamp.Error.Printf("could not send message: `%s`", err)
          }
          resp := <-respChan
          if !bytes.HasPrefix(resp.Bytes(), expBytes) {
            scamp.Error.Printf("`%s` vs `%s`", string(resp.Bytes()), string(expBytes))
            panic("what")
          }
          // scamp.Info.Printf("resp: %s", resp.Bytes())
          wg.Done()
        }()
      }
      wg.Done()
    }()
  }

  wg.Wait()
}