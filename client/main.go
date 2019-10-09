package main

import (
	"fmt"
	"os"
	"protected-notebook/client/handlers"
)

func main() {
	resp, err := handlers.SendPublicKey()
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
