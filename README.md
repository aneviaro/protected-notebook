# Protected Notebook

Protected Notebook is a GoLang study-project to exchange protected file content using the IDEA-CFB algorithm and RSA algorithm.
## Description
### Client-side: 
1) Generates RSA KeyPair and sends the public key to the server-side. Stores private key inside rsa variable.
2) Picks up a file name and send a request to the server.
3) Decrypts IDEA key with RSA private key.
4) Decrypts file content with IDEA key.
### Server-side:
1) Stores RSA public key inside rsa variable.
2) Sends a list of all available files.
3) Finds the file and encrypts file content using randomly generated IDEA key.
4) Encrypt IDEA key with RSA public key.
## Launch

```bash
cd server1
go run main.go
```
```bash
cd client
go run main.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
