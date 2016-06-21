package common

import (
  "fmt"
)

type VerifyRequest struct {
  Version int `json:"version"`
  TargetUrl string `json:"target_url"`
  SupportedActions []string `json:"supported_actions"`
}

func NewVerifyRequest() (VerifyRequest) {
  return VerifyRequest{}
}

func (vr VerifyRequest) IsValid() (err error) {
  if vr.Version == 0 {
    err = fmt.Errorf("must set version")
  } else if vr.TargetUrl == "" {
    err = fmt.Errorf("must set target url")
  } else if len(vr.SupportedActions) == 0 {
    err = fmt.Errorf("must set supported action list")
  }

  return
}

type VerifyResponse struct {
  Status string `json:"status"`
  Message string `json:"message"`
  ActionResults []ActionResult `json:"action_results"`
}

func NewVerifyResponse() (verResp VerifyResponse) {
  return VerifyResponse {
    ActionResults: make([]ActionResult,0),
  }
}

type ActionResult struct {
  Status string `json:"status"`
  Message string `json:"message"`
  Action string `json:"action"`
  Version int `json:"version"`
  TargetUrl string `json:"target_url"`
}