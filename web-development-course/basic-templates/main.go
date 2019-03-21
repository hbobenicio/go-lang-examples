package main

import (
	"html/template"
	"log"
	"os"
)

type templateData struct {
	MSG string
}

func main() {
	tpl := template.Must(template.ParseFiles("index.gohtml"))

	data := templateData{
		MSG: "Hello, Template",
	}

	if err := tpl.Execute(os.Stdout, data); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
