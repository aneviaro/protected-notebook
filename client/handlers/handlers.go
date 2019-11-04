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

	"github.com/rs/xid"
)

var (
	client        *http.Client
	currentClient Client
)

func init() {
	client = &http.Client{}
}

func Login(username string, password string)( *http.Response, error){
	c:=Credentials{
		Password: password,
		Username: username,
	}
	credentialsJSON:=credentialsToJson(c)
	return client.Post("http://localhost:8080/login", "json", strings.NewReader(string(credentialsJSON)))

}
//Client stores current client info
type Client struct {
	PublicKey  *rsa.PublicKey `json:"key"`
	ClientName string         `json:"name"`
}

type Credentials struct {
	Password string `json: "password"`
	Username string `json: "username"`
}

func newClient(key *rsa.PublicKey, name string) Client {
	return Client{
		PublicKey:  key,
		ClientName: name}
}

func clientToJSON(client Client) []byte {
	clientJSON, err := json.Marshal(client)
	if err != nil {
		panic(err)
	}
	return clientJSON
}

func credentialsToJson(credentials Credentials)[]byte{
	credentialsJSON, err := json.Marshal(credentials)
	if err != nil {
		panic(err)
	}
	return credentialsJSON
}

//SendPublicKey to server with client id
func SendPublicKey() ( *http.Response, error) {
	createRSAKeyPair()
	publicKey := rsa_initial.GetPublicKey()
	currentClient = newClient(publicKey, xid.New().String())
	clientJSON := clientToJSON(currentClient)
	fmt.Println("Sending publicKey...")
	return client.Post("http://localhost:8080/rsa", "json", strings.NewReader(string(clientJSON)))
}

//GetFileContent send a request to get content
func GetFileContent(name string) (*http.Response, error) {
	fmt.Println("Getting encrypted file content...")
	return client.Get("http://localhost:8080/file/" + name + "/" + currentClient.ClientName)
}

//DecryptContent that passed from server
func DecryptContent(resp *http.Response) string {
	fmt.Println("Decrypting file content...")
	response := convertHTTPBodyToValidResponce(resp.Body)
	return decryptFileContent(response)
}

func decryptFileContent(resp Resp) string {
	key := rsa_initial.DecryptText(resp.Key)
	content := idea.CFBDecrypter(string(resp.Content), string(key))
	return content
}

//Resp for parsing server responce
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

//GetListOfFile from responce Body
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
