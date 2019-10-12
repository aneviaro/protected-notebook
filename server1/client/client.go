package client

import (
	"crypto/rsa"
	"errors"
	"sync"
)

var (
	clientList []Client
	once       sync.Once
	mtx        sync.Mutex
)

func init() {
	clientList = []Client{}
}

//Client to store clients public Keys
type Client struct {
	PublicKey  *rsa.PublicKey `json:"key"`
	ClientName string         `json:"name"`
}

//AddNewClient setting rsa public key and client
func AddNewClient(client Client) {
	mtx.Lock()
	clientList = append(clientList, client)
	mtx.Unlock()
}

func newClient(key *rsa.PublicKey, name string) Client {
	return Client{
		PublicKey:  key,
		ClientName: name}
}

//GetPublicKey by name
func GetPublicKey(name string) *rsa.PublicKey {
	clientObj, err := findClientByClientName(name)
	if err != nil {
		panic(err)
	}
	return clientObj.PublicKey
}

func findClientByClientName(name string) (Client, error) {
	for _, clientItem := range clientList {
		if clientItem.ClientName == name {
			return clientItem, nil
		}
	}
	return Client{}, errors.New("no such client")
}
