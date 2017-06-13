package verify

import (
	"fmt"
	"os"
	u "os/user"

	"crypto/rand"
	"encoding/base64"

	"io/ioutil"

	"log"
	fp "path/filepath"
)

type AuthTokenStorage struct {
	Path string
}

func NewAuthTokenStorage() (ts *AuthTokenStorage, err error) {
	//TODO: os/user is not supported for cross compiling, find diff way of doing this
	user, err := u.Current()
	if err != nil {
		return
	}

	return &AuthTokenStorage{
		Path: fp.Join(user.HomeDir, "retailops"),
	}, nil
}

func (ats *AuthTokenStorage) CreateDirectoryIfMissing() (err error) {
	_, err = os.Open(ats.Path)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok && pathErr.Err.Error() == "no such file or directory" {
			return ats.doFolderCreate()
		}
		return

	}

	return
}

func (ats *AuthTokenStorage) OverwriteIntegrationAuthKey(token string) (err error) {
	tokenPath := fp.Join(ats.Path, "integration_auth_token")
	var file *os.File
	if _, err = os.Stat(tokenPath); os.IsNotExist(err) {
		file, err = os.Create(tokenPath)
		if err != nil {
			return
		}
	} else {
		file, err = os.OpenFile(tokenPath, os.O_WRONLY|os.O_TRUNC, 0)
		if err != nil {
			return
		}
	}

	n, err := file.WriteString(token)
	if err != nil {
		return
	}

	if n != len(token) {
		log.Fatal("huh")
	}

	return
}

func (ats *AuthTokenStorage) GenerateIntegrationAuthToken() (err error) {
	tokenPath := fp.Join(ats.Path, "integration_auth_token")
	fmt.Println("Token path: ", tokenPath)

	token, err := randomToken()
	if err != nil {
		return
	}

	file, err := os.Create(tokenPath)
	if err != nil {
		fmt.Println("Could not create folder: ", err)
		return
	}

	_, err = file.WriteString(token)
	if err != nil {
		fmt.Println("Could not create file: ", err)
	}

	return
}

func (ats *AuthTokenStorage) ReadToken() (token string, err error) {
	tokenPath := fp.Join(ats.Path, "integration_auth_token")
	file, err := os.Open(tokenPath)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	token = string(buf)
	return
}

func randomToken() (token string, err error) {
	randBytes := make([]byte, 24, 24)
	read, err := rand.Read(randBytes)
	if err != nil {
		err = fmt.Errorf("could not generate all rand bytes needed. only read %d of 18", read)
		return
	}
	token = base64.StdEncoding.EncodeToString(randBytes)

	return
}

func (ats *AuthTokenStorage) doFolderCreate() (err error) {
	fmt.Println("Creating folder: ", ats.Path)
	err = os.Mkdir(ats.Path, 0700)
	return
}
