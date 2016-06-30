package common

import (
  "fmt"
  "strings"
  "bytes"
  "text/template"
)

type VerifyRequest struct {
  ApiKey string `json:"apikey,omitempty"`
  IntegrationAuthKey string `json:"integration_auth_key"`
  Version int `json:"version"`
  TargetUrl string `json:"target_url"`
  SupportedActions []string `json:"supported_actions"`
  IntegrationName string `json:"integration_name"`
}

func NewVerifyRequest() (*VerifyRequest) {
  return &VerifyRequest{}
}

func (vr VerifyRequest) IsValid() (err error) {
  if vr.Version == 0 {
    err = fmt.Errorf("must set version")
  } else if vr.TargetUrl == "" {
    err = fmt.Errorf("must set target url")
  } else if len(vr.SupportedActions) == 0 {
    err = fmt.Errorf("must set supported action list")
  } else if len(vr.IntegrationName) == 0 {
    err = fmt.Errorf("must provide integration name")
  }

  return
}

type VerifyResponse struct {
  Status string `json:"status"`
  Message string `json:"message"`
  ActionResults []ActionResult `json:"action_results"`

  // The following can be injected by the API dispatcher
  Error string `json:"ERROR"`
  ErrorCode string `json:"ERRORCODE"`
}

func NewVerifyResponse() (verResp VerifyResponse) {
  return VerifyResponse {
    ActionResults: make([]ActionResult,0),
  }
}

func (verResp VerifyResponse) NiceError() (errMsg string) {
  var buf bytes.Buffer

  tmpl,err := template.New("T").Parse(`{{.FailCount}} of {{len .ActionResults}} actions failed
{{range .ActionResults}}{{.Status}}: {{.TargetUrl}} {{.Message}}
{{end}}
`)
  if err != nil {
    panic(err.Error())
  }


  err = tmpl.Execute(&buf, verResp)
  if err != nil {
    panic(err.Error())
  }

  return string(buf.Bytes())
}

func (verResp VerifyResponse) FailCount() (count int) {
  count = 0
  for _,actionResult := range verResp.ActionResults {
    if actionResult.Status == "error" {
      count += 1
    }
  }

  return count
}

type ActionResult struct {
  Status string `json:"status"`
  Message string `json:"message"`
  Action string `json:"action"`
  Version int `json:"version"`
  TargetUrl string `json:"target_url"`
}

type RegistrationRequest struct {
  Name string `json:"name"`
  Handle string `json:"handle"`
  Interactions []RegistrationInteraction `json:"interactions"`
}

func NewRegistrationRequest(name, authKey string) (req *RegistrationRequest) {
  noSpaces := strings.Replace(name, " ", "_", -1)
  handleBase := strings.ToUpper(noSpaces)
  handle := fmt.Sprintf("%s_%s",handleBase,authKey)

  return &RegistrationRequest {
    Name: name,
    Handle: handle,
    Interactions: make([]RegistrationInteraction,0),
  }
}

func (req *RegistrationRequest)AddInteraction(action, endpointUrl string) {
  req.Interactions = append(req.Interactions, RegistrationInteraction {
    Action: action,
    EndpointUrl: endpointUrl,
  })
}

type RegistrationInteraction struct {
  Action string `json:"action"`
  EndpointUrl string `json:"endpoint_url"`
}