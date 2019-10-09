package idea

import (
	"crypto/cipher"
)

func CFBDecrypter(ciphertxt string, keyParam string) string {
	key := []byte(keyParam)
	ciphertext := []byte(ciphertxt)
	block, err := NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < block.BlockSize() {
		panic("Ciphertext too short")
	}

	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}
