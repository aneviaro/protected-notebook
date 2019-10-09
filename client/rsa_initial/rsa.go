package rsa_initial

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
)

var (
	privateKey *rsa.PrivateKey
)

func GenerateKeyPair() error {
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	return nil
}

func GetPublicKey() *rsa.PublicKey {
	return &privateKey.PublicKey
}

func DecryptText(ciphertext []byte) []byte {
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		fmt.Println("Something happened while decrypting IDEA Key with RSA private key")
	}
	return plaintext
}
