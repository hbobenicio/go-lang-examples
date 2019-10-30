package util

import (
	"log"
	"net/http"
)

// WriteAndLogErr Writes and logs any errors that may eventually occur
func WriteAndLogErr(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Println("write failed:", err)
	}
}
