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
	defer resp.Body.Close()
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
	if handlers.IsAdmin() {
		showManageAccess()
	} else {
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
			if resp.StatusCode == 200 {
				decryptedContent := handlers.DecryptContent(resp)
				fmt.Printf("%s\n", decryptedContent)
			} else {
				fmt.Println(resp.Status)
				continue
			}
		}
	}
}
func showManageAccess() {
	var filename, username string
	fmt.Println("Grant access")
	fmt.Println("filename:")
	fmt.Fscan(os.Stdin, &filename)
	fmt.Println("username:")
	fmt.Fscan(os.Stdin, &username)
	resp, err := handlers.SendGrantAccessRequest(filename, username)
	if err != nil {
		fmt.Println("Something went wrong while sending grant request")
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		showManageAccess()
	} else {
		fmt.Println("Successfully granted")
		showManageAccess()
	}
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
