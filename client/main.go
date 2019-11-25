package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"protected-notebook/client/handlers"
)

func main() {
	resp, err := doLogin()
	if err != nil {
		fmt.Println("Something went wrong while sending RSA")
		panic(err)
	}
	for resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		resp, err = doLogin()
		if err != nil {
			fmt.Println("Something went wrong while sending RSA")
			panic(err)
		}
	}
	files := handlers.GetListOfFile(resp)

	for {
		fmt.Println("Choose a file: ")
		for _, fileItem := range files {
			fmt.Printf("%v\n", fileItem.Name)
		}
		var name string
		fmt.Fscan(os.Stdin, &name)
		resp, err = handlers.GetFileContent(name)
		if err != nil {
			fmt.Println("Something went wrong while getting file content")
			panic(err)
		}
		decryptedContent := handlers.DecryptContent(resp)
		fmt.Printf("%s\n", decryptedContent)
	}
	defer resp.Body.Close()
}
func doLogin() (*http.Response, error) {
	var username, password string
	fmt.Println("username:")
	fmt.Fscan(os.Stdin, &username)
	fmt.Println("password:")
	fmt.Fscan(os.Stdin, &password)
	h := md5.New()
	h.Write([]byte(password))
	pw1 := hex.EncodeToString(h.Sum(nil))
	h.Reset()
	h.Write([]byte(pw1))
	handlers.SetCredentials(username, h.Sum(nil))
	return handlers.SendPublicKey()
}
