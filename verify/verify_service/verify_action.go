package verify_service

import (
  "bytes"
  "fmt"

  "encoding/json"

  "os"
  "io/ioutil"
  "time"

  "github.com/gudtech/scamp-go/scamp"
  "github.com/gudtech/retailops-sdk/verify/verify"
  "github.com/gudtech/retailops-sdk/verify/common"

)

var TicketPath string = ""

func VerifyAction(msg *scamp.Message, client *scamp.Client) {
  var err error
  verResp := common.NewVerifyResponse()
  // scamp.Error.Printf("received new verify request size: %d", len(msg.Bytes()))

  // scamp.Error.Printf("msg ticket: %s", msg.GetTicket())

  respMsg := scamp.NewResponseMessage()
  respMsg.SetRequestId(msg.RequestId)

  var req = common.NewVerifyRequest()
  err = json.NewDecoder(bytes.NewReader(msg.Bytes())).Decode(&req)
  if err != nil {
    verResp.Status = "error"
    verResp.Message = "request is not valid JSON"
    respMsg.WriteJson(verResp)
    _,_ = client.Send(respMsg)
    return
  }

  if err = req.IsValid(); err != nil {
    verResp.Status = "error"
    verResp.Message = fmt.Sprintf("request is not valid: `%s`", err.Error())
    respMsg.WriteJson(verResp)
    _,_ = client.Send(respMsg)
    return
  }

  doVerificationRequest(req, &verResp, false)

  if verResp.Status == "success" {
    err = doRegistration(msg.GetTicket(), req, &verResp)
    if err != nil {
      scamp.Error.Printf("error calling registration: %s", err.Error())
      respMsg.WriteJson(map[string]string{
        "error": err.Error(),
      })
      _,err = client.Send(respMsg)
      if err != nil {
        scamp.Error.Printf("err: %s", err.Error())
      }
      return
    }
  }

  respMsg.WriteJson(verResp)
  _,err = client.Send(respMsg)
  if err != nil {
    return
  }
}

func doRegistration(identifyingToken string, req *common.VerifyRequest, resp *common.VerifyResponse) (err error) {
  regReq := common.NewRegistrationRequest(req.IntegrationName, req.IntegrationAuthKey)
  for _,actionResult := range resp.ActionResults {
    regReq.AddInteraction(actionResult.Action, actionResult.TargetUrl)
  }

  stationFile,err := os.Open(TicketPath)
  if err != nil {
    return
  }
  stationBytes,err := ioutil.ReadAll(stationFile)
  if err != nil {
    return
  }

  station := string(stationBytes)

  msg := scamp.NewRequestMessage()
  msg.SetTicket(station)
  msg.SetIdentifyingToken(identifyingToken)
  msg.SetRequestId(1)
  msg.WriteJson(regReq)
  msg.SetTicket(station)

  respchan,err := scamp.MakeJsonRequest("main", "Integration.Channel.register", 1, msg)
  if err != nil {
    return
  }

  select {
  case respMsg := <-respchan:
    strresp := string(respMsg.Bytes())
    scamp.Info.Printf("%s %s %s", respMsg.GetError(), respMsg.GetErrorCode(),   strresp)
    if respMsg.GetError() != "" {
      err = fmt.Errorf(respMsg.GetError())
      return
    }
  case <-time.After(time.Duration(10 * time.Second)):
    scamp.Error.Printf("timeout...")
  }

  return
}

func doVerificationRequest(verReq *common.VerifyRequest, verResp *common.VerifyResponse, mungeAuthToken bool) {
  var failCount int = 0
  for _,action := range verReq.SupportedActions {
    //TODO: abs path
    schemaFile,err := os.Open(fmt.Sprintf("/github.com/gudtech/retailops-sdk/verify/schema/%s_v%d.json", action, verReq.Version))
    if err != nil {
      verResp.ActionResults = append(verResp.ActionResults, common.ActionResult {
        Status: "error",
        Message: err.Error(),
        Action: action,
        Version: verReq.Version,
        TargetUrl: "",
      })

      failCount += 1
      continue
    }

    exampleFile,err := os.Open(fmt.Sprintf("/github.com/gudtech/retailops-sdk/verify/tests/%s_v%d_ex_1.json", action, verReq.Version))
    if err != nil {
      verResp.ActionResults = append(verResp.ActionResults, common.ActionResult {
        Status: "error",
        Message: err.Error(),
        Action: action,
        Version: verReq.Version,
        TargetUrl: "",
      })

      failCount += 1
      continue
    }

    var integrationAuthKey = verReq.IntegrationAuthKey
    if mungeAuthToken {
      integrationAuthKey = fmt.Sprintf("!!%s!!",verReq.IntegrationAuthKey)
    }

    /*
      Make the normal request, proper settings, and everything
    */
    err = verify.Request(verReq.TargetUrl, integrationAuthKey, schemaFile, exampleFile, true, 200)//TODO double check 200
    if err != nil {
      verResp.ActionResults = append(verResp.ActionResults, common.ActionResult {
        Status: "error",
        Message: err.Error(),
        Action: action,
        Version: verReq.Version,
        TargetUrl: fmt.Sprintf("%s/%s", verReq.TargetUrl, action),
      })

      failCount += 1
      continue
    }

    // Check that they reject requests with bad integration_auth_tokens
    _,err = schemaFile.Seek(0,0)
    if err != nil { return }
    _,err = exampleFile.Seek(0,0)
    if err != nil { return }
    err = verify.Request(verReq.TargetUrl, fmt.Sprintf("!!!%s", integrationAuthKey), schemaFile, exampleFile, true, 200) //TODO check 200
    if err == nil {
      verResp.ActionResults = append(verResp.ActionResults, common.ActionResult {
        Status: "error",
        Message: "failed to check integration_auth_token. expected HTTP 401", //NOTE: may need to allow 401 in certain cases
        Action: action,
        Version: verReq.Version,
        TargetUrl: fmt.Sprintf("%s/%s", verReq.TargetUrl, action),
      })

      failCount += 1
      continue
    }

    verResp.ActionResults = append(verResp.ActionResults, common.ActionResult {
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
    verResp.Message = verResp.NiceError() // fmt.Sprintf("%d verification requests failed", failCount)
  }

  return
}
