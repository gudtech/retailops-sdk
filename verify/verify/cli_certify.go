package verify

import (
  "fmt"
  "bytes"
  "encoding/json"

  "net/http"
  "time"

  // "io/ioutil"

  "github.com/gudtech/retailops-sdk/verify/common"
)

var certifyClient = &http.Client{
  Timeout: time.Minute * 5,
}

func doCertify(cliExec CLIExecution) (err error) {
  if len(cliExec.CertifyActions) == 0 {
    err = fmt.Errorf("must set at least 1 action to certify")
    return
  } else if len(cliExec.IntegrationAuthKey) == 0 {
    err = fmt.Errorf("integration auth key is not set. please refer to documentation for how to generate auth key")
    return
  }

  var verReq = common.VerifyRequest {
    Version: 1,
    TargetUrl: cliExec.BaseURL,
    SupportedActions: cliExec.CertifyActions,
    IntegrationAuthKey: cliExec.IntegrationAuthKey,
    IntegrationName: cliExec.IntegrationName,
  }

  var buf bytes.Buffer
  err = json.NewEncoder(&buf).Encode(verReq)
  if err != nil {
    return
  }

  // url := fmt.Sprintf("%s?apikey=%s",cliExec.ROCertifyURL,cliExec.ApiKey)
  url := fmt.Sprintf("%s?apikey=%s","https://api.retailops.com/integration/channel/certify.json",cliExec.ApiKey);
  resp,err := certifyClient.Post(url, "application/json", &buf)
  defer resp.Body.Close()
  if err != nil {
    return
  }

  // respBuf,err := ioutil.ReadAll(resp.Body)
  // if err != nil {
  //   return
  // }
  // panic(string(respBuf))

  var apiResp common.VerifyResponse
  err = json.NewDecoder(resp.Body).Decode(&apiResp)
  if err != nil {
    return
  }

  if apiResp.Status == "error" {
    err = fmt.Errorf("certification failed:\n%s", apiResp.Message)
    return
  } else if apiResp.Error != "" {
    err = fmt.Errorf("certification failed: %s/%s", apiResp.Error, apiResp.ErrorCode)
    return
  }

  return

}
