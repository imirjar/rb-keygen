package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting pwd", err)
		return
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Приватный ключ в формате PEM
	privateKeyPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateFile, err := os.Create(fmt.Sprintf(pwd + "/private_key.pem"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer privateFile.Close()
	pem.Encode(privateFile, &privateKeyPEM)

	// Публичный ключ в формате PEM
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)

	publicKeyPEM := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicFile, err := os.Create(fmt.Sprintf(pwd + "/public_key.pem"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer publicFile.Close()
	pem.Encode(publicFile, &publicKeyPEM)

	fmt.Println("Keys generated and saved to 'private_key.pem' and 'public_key.pem'")
}
