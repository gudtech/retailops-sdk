package verify_service

type VerifyRequest struct {
  Version int `json:"version"`
  TargetUrl string `json:"target_url"`
  SupportedActions []string `json:"supported_actions"`
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