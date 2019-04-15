package main

import (
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	pemFile, err := os.Open("GlobalSignRootCA-R2.crt.pem")
	if err != nil {
		log.Fatalf("não foi possível abrir o arquivo: %v", err)
	}
	defer pemFile.Close()

	pemData, err := ioutil.ReadAll(pemFile) // could be improved to use bufio.NewReader
	if err != nil {
		log.Fatalf("não foi possível ler o arquivo: %v", err)
	}

	// Vários blocos de certificados e/ou chaves públicas podem ser obtidos a partir daqui
	i := 0
	for block, rest := pem.Decode(pemData); block != nil; block, rest = pem.Decode(rest) {
		i++
		fmt.Println("Bloco", i)

		switch block.Type {
		case "CERTIFICATE":

			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				log.Fatalf("não foi possível fazer parsing do certificado: %v", err)
			}

			fmt.Println("Issuer:", cert.Issuer)
			fmt.Println("Subject:", cert.Subject)
			fmt.Println("Public Key Algorithm:", cert.PublicKeyAlgorithm)
			fmt.Printf("Public Key Go Type: %T\n", cert.PublicKey)
			// fmt.Printf("%+v\n", cert)

			switch cert.PublicKey.(type) {
			case *rsa.PublicKey:
				pubKey := cert.PublicKey.(*rsa.PublicKey)
				fmt.Printf("Chave Pública: %+v\n", pubKey)

			case *dsa.PublicKey:
				pubKey := cert.PublicKey.(*dsa.PublicKey)
				fmt.Printf("Chave Pública: %+v\n", pubKey)

			case *ecdsa.PublicKey:
				pubKey := cert.PublicKey.(*ecdsa.PublicKey)
				fmt.Printf("Chave Pública: %+v\n", pubKey)

			default:
				log.Fatalln("Tipo de chave pública não suportado")
			}

			// pubKey, err := x509.ParsePKIXPublicKey(pubKey)
			// if err != nil {
			// 	log.Fatalf("não foi possível fazer parsing dos campos x509 do bloco pem: %v", err)
			// }

		case "PUBLIC KEY":
			pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				log.Fatalf("não foi possível fazer parsing dos campos x509 do bloco pem: %v", err)
			}

			fmt.Printf("Got a %T, with remaining data: %q", pubKey, rest)

		default:
			log.Fatalln("tipo de bloco não suportado:", block.Type)
		}

	}
}
