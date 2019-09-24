package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const cookieSessionKey = "session-id"

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieSessionKey)
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  cookieSessionKey,
			Value: id.String(),
			Path:  "/",
			// Secure: true,
			// HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	fmt.Println(cookie)
	io.WriteString(w, cookie.String())
}
