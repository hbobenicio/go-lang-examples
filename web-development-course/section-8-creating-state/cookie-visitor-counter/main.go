package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const visitCounterCookieKey = "visit-counter"

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	visitCounterCookie, err := r.Cookie(visitCounterCookieKey)
	if err == http.ErrNoCookie {
		log.Println("creating a new visit counter cookie for you")

		visitCounterCookie = &http.Cookie{
			Name:   visitCounterCookieKey,
			Value:  "0",
			Domain: "localhost",
			Path:   "/",
			MaxAge: 3600,
		}

	} else if err != nil {
		errMsg := fmt.Sprintf("unexpected error while getting cookie from request: %v", err)
		fmt.Fprintln(os.Stderr, errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Parsing the old vist count value
	visitCount, err := strconv.Atoi(visitCounterCookie.Value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing cookie: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	visitCount++
	visitCounterCookie.Value = strconv.Itoa(visitCount)

	http.SetCookie(w, visitCounterCookie)
	fmt.Fprintf(w, "Visit Counter: %s\n", visitCounterCookie.Value)
}
