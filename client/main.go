package main

import (
	"fmt"
	"net/http"
	"os"
	"protected-notebook/client/handlers"
)

func main() {
	resp, err := doLogin()
	if err != nil {
		fmt.Println("Something went wrong while authentication")
		panic(err)
	}
	for resp.StatusCode != 200 {
		fmt.Println(http.StatusText(resp.StatusCode))
		resp,err=doLogin()
	}

	resp, err = handlers.SendPublicKey()
	if err != nil {
		fmt.Println("Something went wrong while sending RSA")
		panic(err)
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
	return handlers.Login(username, password)
}
