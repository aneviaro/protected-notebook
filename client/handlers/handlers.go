package handlers

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"protected-notebook/client/file"
	"protected-notebook/client/idea"
	"protected-notebook/client/rsa_initial"
	"strings"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}

func SendPublicKey() (*http.Response, error) {
	createRSAKeyPair()
	publicKey := rsa_initial.GetPublicKey()
	pubInJSON := publicKeyToJSON(publicKey)
	fmt.Println("Sending publicKey...")
	return client.Post("http://localhost:8080/rsa", "json", strings.NewReader(string(pubInJSON)))
}

func GetFileContent(name string) (*http.Response, error) {
	fmt.Println("Getting encrypted file content...")
	return client.Get("http://localhost:8080/file/" + name)
}

func DecryptContent(resp *http.Response) string {
	fmt.Println("Decrypting file content...")
	responce := convertHTTPBodyToValidResponce(resp.Body)
	return decryptFileContent(responce)
}

func decryptFileContent(resp Resp) string {
	key := rsa_initial.DecryptText(resp.Key)
	content := idea.CFBDecrypter(string(resp.Content), string(key))
	return content
}

type Resp struct {
	Content []byte `json:"content"`
	Key     []byte `json:"key"`
}

func convertHTTPBodyToValidResponce(httpBody io.ReadCloser) Resp {

	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		fmt.Println("Something went wrong while reading server responce")
		panic(err)
	}

	var resp Resp
	json.Unmarshal(body, &resp)
	return resp
}

func GetListOfFile(resp *http.Response) []file.File {
	files, err := convertHTTPBodyToListOfFiles(resp.Body)
	if err != nil {
		fmt.Println("Something went wrong while getting list of files")
		panic(err)
	}
	return files
}

func createRSAKeyPair() {
	fmt.Println("Generating a keypair...")
	err := rsa_initial.GenerateKeyPair()
	if err != nil {
		panic(err)
	}
}

func publicKeyToJSON(publicKey *rsa.PublicKey) []byte {
	pubInJason, err := json.Marshal(publicKey)
	if err != nil {
		panic(err)
	}
	return pubInJason
}

func convertHTTPBodyToListOfFiles(httpBody io.ReadCloser) ([]file.File, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return nil, err
	}

	var files []file.File
	json.Unmarshal(body, &files)
	return files, nil
}