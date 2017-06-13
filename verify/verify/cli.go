package verify

import (
	"fmt"
)

type SchemaExample struct {
	SchemaPath  string
	ExamplePath string
}

type CLIExecution struct {
	Action       string
	SchemaPath   string
	ExamplesPath string
	BaseURL      string
	SchemaFilter string

	ApiKey             string
	IntegrationAuthKey string
	IntegrationName    string

	SchemaPathIsDir   bool
	ExamplesPathIsDir bool
	StopOnError       bool
	Verbose           bool
	CertifyActions    []string
	UseAuthKey        bool
	// ROCertifyURL string
}

var HR string = "----------------"

func Execute(cliExec CLIExecution) (err error) {
	if cliExec.Action == "test" {
		return doLocalVerify(cliExec)
	} else if cliExec.Action == "certify" {
		return doCertify(cliExec)
	} else if cliExec.Action == "generate_token" {
		return doGenerateToken(cliExec)
	} else if cliExec.Action == "show_token" {
		return doTokenShow(cliExec)
	} else if cliExec.Action == "install_token" {
		return doInstallIntegrationAuthKey(cliExec)
	}
	err = fmt.Errorf("unhandled action %s", cliExec.Action)
	return
}
