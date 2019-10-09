package rsa_initial

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"sync"
)

var (
	publicKey *rsa.PublicKey
	once      sync.Once
)

//SetRSA setting rsa public key
func SetRSA(public *rsa.PublicKey) {
	once.Do(func() {
		publicKey = public
	})
}

//EncryptText encrypt text in []byte with sha1
func EncryptText(text []byte) []byte {
	sha1 := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(sha1, rand.Reader, publicKey, text, nil)
	if err != nil {
		panic(err)
	}
	return ciphertext
}
