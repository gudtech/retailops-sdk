package verify

import (
	"fmt"
)

func doInstallIntegrationAuthKey(cliExec CLIExecution) (err error) {
	if len(cliExec.IntegrationAuthKey) == 0 {
		//TODO: check format of usage below
		err = fmt.Errorf("usage: `verify install-auth-key INTEGRATION_AUTH_KEY`")
		return
	}

	ats, err := NewAuthTokenStorage()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("cliExec.IntegrationAuthKey: ", cliExec.IntegrationAuthKey)
	err = ats.OverwriteIntegrationAuthKey(cliExec.IntegrationAuthKey)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	return
}

func doGenerateToken(cliExec CLIExecution) (err error) {
	fmt.Println("Generating token...")
	ats, err := NewAuthTokenStorage()
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
	ats, err := NewAuthTokenStorage()
	if err != nil {
		return
	}

	token, err := ats.ReadToken()
	if err != nil {
		return
	}

	fmt.Println("==== INTEGRATION AUTH KEY BELOW ====")
	fmt.Println(token)
	fmt.Println("==== INTEGRATION AUTH KEY DONE ====")

	return
}
