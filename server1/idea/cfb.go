package idea

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

//CFBEncrypter encrypt plaintext in with IDEA-CFB
func CFBEncrypter(plaintext []byte) ([]byte, []byte) {
	hexKey, err := randomHex(16)
	if err != nil {
		panic(err)
	}
	key, _ := hex.DecodeString(hexKey)
	block, err := NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, block.BlockSize()+len(plaintext))
	iv := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[block.BlockSize():], plaintext)
	return key, ciphertext
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
