package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hbobenicio/go-lang-examples/x509-certs-pem/config"
	mid "github.com/hbobenicio/go-lang-examples/x509-certs-pem/middleware"
)

func main() {
	cfg := config.New()
	cfg.LoadFromEnv()

	router := mux.NewRouter()
	router.Handle("/certs/info", certsInfoPipeline()).Methods("POST")
	router.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("servidor escutando na porta", cfg.ServerPort)
	log.Fatalln(http.ListenAndServe(":"+cfg.ServerPort, router))
}

func certsInfoPipeline() http.Handler {
	return mid.LogRequest(http.HandlerFunc(certInfoHandler))
}

func certInfoHandler(w http.ResponseWriter, r *http.Request) {
	type RequestPayload struct {
		Domain string `json:"domain"`
	}

	type ResponsePayload struct {
		Issuer    string `json:"issuer"`
		Subject   string `json:"subject"`
		NotBefore string `json:"notBefore"`
		NotAfter  string `json:"notAfter"`
	}

	var payload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	addr := fmt.Sprintf("%s:443", payload.Domain)

	conn, err := tls.Dial("tcp", addr, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println("erro de conexão:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	if err := conn.Handshake(); err != nil {
		log.Println("erro de handshake:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var response []ResponsePayload
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		response = append(response, ResponsePayload{
			Issuer:    cert.Issuer.String(),
			Subject:   cert.Subject.String(),
			NotBefore: cert.NotAfter.String(),
			NotAfter:  cert.NotAfter.String(),
		})
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("error: certs info: encoding json response:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// pemFile, err := os.Open("certificates.pem")
// if err != nil {
// 	log.Fatalf("não foi possível abrir o arquivo: %v", err)
// }
// defer pemFile.Close()

// pemData, err := ioutil.ReadAll(pemFile) // could be improved to use bufio.NewReader
// if err != nil {
// 	log.Fatalf("não foi possível ler o arquivo: %v", err)
// }

// // Vários blocos de certificados e/ou chaves públicas podem ser obtidos a partir daqui
// i := 0
// for block, rest := pem.Decode(pemData); block != nil; block, rest = pem.Decode(rest) {
// 	i++
// 	fmt.Println("Bloco", i)

// 	switch block.Type {
// 	case "CERTIFICATE":

// 		cert, err := x509.ParseCertificate(block.Bytes)
// 		if err != nil {
// 			log.Fatalf("não foi possível fazer parsing do certificado: %v", err)
// 		}

// 		fmt.Println("  Issuer:", cert.Issuer)
// 		fmt.Println("  Subject:", cert.Subject)
// 		fmt.Println("  Public Key Algorithm:", cert.PublicKeyAlgorithm)
// 		fmt.Printf("  Public Key Go Type: %T\n", cert.PublicKey)
// 		// fmt.Printf("%+v\n", cert)

// 		switch cert.PublicKey.(type) {
// 		case *rsa.PublicKey:
// 			pubKey := cert.PublicKey.(*rsa.PublicKey)
// 			fmt.Printf("  Chave Pública: %+v\n", pubKey)

// 		case *dsa.PublicKey:
// 			pubKey := cert.PublicKey.(*dsa.PublicKey)
// 			fmt.Printf("  Chave Pública: %+v\n", pubKey)

// 		case *ecdsa.PublicKey:
// 			pubKey := cert.PublicKey.(*ecdsa.PublicKey)
// 			fmt.Printf("  Chave Pública: %+v\n", pubKey)

// 		default:
// 			log.Fatalln("Tipo de chave pública não suportado")
// 		}

// 		// pubKey, err := x509.ParsePKIXPublicKey(pubKey)
// 		// if err != nil {
// 		// 	log.Fatalf("não foi possível fazer parsing dos campos x509 do bloco pem: %v", err)
// 		// }

// 	case "PUBLIC KEY":
// 		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
// 		if err != nil {
// 			log.Fatalf("não foi possível fazer parsing dos campos x509 do bloco pem: %v", err)
// 		}

// 		fmt.Printf("Got a %T, with remaining data: %q", pubKey, rest)

// 	default:
// 		log.Fatalln("tipo de bloco não suportado:", block.Type)
// 	}
// 	fmt.Println()
// }
