package verify_service

import (
  "bytes"
  "fmt"

  "encoding/json"

  "os"

  "github.com/gudtech/scamp-go/scamp"
  "github.com/gudtech/retailops-sdk/verify/verify"
)

// var httpClient = &http.Client{
//   Timeout: time.Second * 15,
// }

func (vr VerifyRequest) IsValid() (reason err) {
  if !vr.Version > 0 {
    err = fmt.Errorf("must set version")
  } else if vr.TargetUrl == "" {
    err = fmt.Errorf("must set target url")
  } else if len(vr.SupportedActions) == 0 {
    err = fmt.Errorf("must set supported action list")
  }

  return
}

func VerifyAction(msg *scamp.Message, client *scamp.Client) {
  var err error
  verResp := NewVerifyResponse()
  // scamp.Error.Printf("received new verify request size: %d", len(msg.Bytes()))

  respMsg := scamp.NewResponseMessage()
  respMsg.SetRequestId(msg.RequestId)

  var req VerifyRequest
  err = json.NewDecoder(bytes.NewReader(msg.Bytes())).Decode(&req)
  if err != nil {
    verResp.Status = "error"
    verResp.Message = "request is not valid JSON"
    respMsg.WriteJson(verResp)
    _,_ = client.Send(respMsg)
    return
  }

  if !req.IsValid() {
    verResp.Status = "error"
    verResp.Message = "request is not valid"
    respMsg.WriteJson(verResp)
    _,_ = client.Send(respMsg)
    return
  }

  doVerificationRequest(req, &verResp)

  if verResp.Status == "success" {
    doRegistration()
  }

  respMsg.WriteJson(verResp)
  _,err = client.Send(respMsg)
  if err != nil {
    return
  }
}


func doVerificationRequest(verReq VerifyRequest, verResp *VerifyResponse) {
  // url,err := url.Parse(verReq.TargetUrl)
  // if err != nil {
  //   verResp.Status = "error"
  //   verResp.Message = fmt.Sprintf("could not parse target_url: %s", err.Error())
  //   return
  // }

  // originalPath := url.Path
  // if originalPath[len(originalPath)-1] != '/' {
  //   originalPath = fmt.Sprintf("%s/", originalPath)
  // }

  var failCount int = 0

  for _,action := range verReq.SupportedActions {
    schemaFile,err := os.Open(fmt.Sprintf("/Users/xavierlange/code/gudtech/workspace/src/github.com/gudtech/retailops-sdk/schema/schemata/%s_v%d.json", action, verReq.Version))
    if err != nil {
      verResp.ActionResults = append(verResp.ActionResults, ActionResult {
        Status: "error",
        Message: err.Error(),
        Action: action,
        Version: verReq.Version,
        TargetUrl: "",
      })

      failCount += 1
      continue
    }

    exampleFile,err := os.Open(fmt.Sprintf("/Users/xavierlange/code/gudtech/workspace/src/github.com/gudtech/retailops-sdk/schema/schemata/%s_v%d_ex_1.json", action, verReq.Version))
    if err != nil {
      verResp.ActionResults = append(verResp.ActionResults, ActionResult {
        Status: "error",
        Message: err.Error(),
        Action: action,
        Version: verReq.Version,
        TargetUrl: "",
      })

      failCount += 1
      continue
    }

    err = verify.Request(verReq.TargetUrl, schemaFile, exampleFile, true)
    if err != nil {
      verResp.ActionResults = append(verResp.ActionResults, ActionResult {
        Status: "error",
        Message: err.Error(),
        Action: action,
        Version: verReq.Version,
        TargetUrl: fmt.Sprintf("%s/%s", verReq.TargetUrl, action),
      })

      failCount += 1
      continue
    }

    // resp,err := httpClient.Get(url.String())
    // if err != nil {
    //   verResp.ActionResults = append(verResp.ActionResults, ActionResult {
    //     Status: "error",
    //     Message: err.Error(),
    //     Action: action,
    //     Version: verReq.Version,
    //     TargetUrl: url.String(),
    //   })

    //   failCount += 1
    //   continue
    // }

    // _,err = ioutil.ReadAll(resp.Body)
    // if err != nil {
    //   verResp.ActionResults = append(verResp.ActionResults, ActionResult {
    //     Status: "error",
    //     Message: err.Error(),
    //     Action: action,
    //     Version: verReq.Version,
    //     TargetUrl: url.String(),
    //   })
    //   failCount += 1
    //   continue
    // }

    verResp.ActionResults = append(verResp.ActionResults, ActionResult {
      Status: "success",
      Message: "",
      Action: action,
      Version: verReq.Version,
      TargetUrl: fmt.Sprintf("%s/%s", verReq.TargetUrl, action),
    })
  }

  if failCount == 0 {
    verResp.Status = "success"
  } else {
    verResp.Status = "error"
    verResp.Message = fmt.Sprintf("%d verification requests failed", failCount)
  }

  return
}