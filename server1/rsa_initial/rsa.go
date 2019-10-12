package rsa_initial

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
)

//EncryptText encrypt text in []byte with sha1
func EncryptText(text []byte, publicKey *rsa.PublicKey) []byte {
	sha1 := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(sha1, rand.Reader, publicKey, text, nil)
	if err != nil {
		panic(err)
	}
	return ciphertext
}
