package verify

import (
  "fmt"
)

func doInstallIntegrationAuthKey(cliExec CLIExecution) (err error) {
  if len(cliExec.IntegrationAuthKey) == 0 {
    err = fmt.Errorf("usage: `verify -integration-auth-key AUTH_KEY install-token`")
    return
  }

  ats,err := NewAuthTokenStorage()
  if err != nil {
    return
  }

  err = ats.OverwriteIntegrationAuthKey(cliExec.IntegrationAuthKey)
  if err != nil {
    return
  }

  return
}



func doGenerateToken(cliExec CLIExecution) (err error) {
  ats,err := NewAuthTokenStorage()
  if err != nil {
    return
  }

  err = ats.CreateDirectoryIfMissing()
  if err != nil {
    return
  }

  err = ats.GenerateIntegrationAuthToken()
  if err != nil {
    return
  }

  return
}

func doTokenShow(cliExec CLIExecution) (err error) {
  ats,err := NewAuthTokenStorage()
  if err != nil {
    return
  }

  token,err := ats.ReadToken()
  if err != nil {
    return
  }

  fmt.Println("==== INTEGRATION AUTH KEY BELOW ====")
  fmt.Println(token)
  fmt.Println("==== INTEGRATION AUTH KEY DONE ====")

  return
}
