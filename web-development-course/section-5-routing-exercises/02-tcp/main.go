package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"text/template"
)

var rootPageTemplate *template.Template
var applyPageTemplate *template.Template

func init() {
	rootPageTemplate = template.Must(template.ParseFiles("templates/pages/index.go.html"))
	applyPageTemplate = template.Must(template.ParseFiles("templates/pages/apply/index.go.html"))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error while accepting new connection: %v\n", err)
			continue
		}

		log.Println("new client connection")
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	log.Println("reading from connection and parsing the request message:")

	var method, URI string
	firstLine := true

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// if its the first line, then this is the Resquest Line of the HTTP Request
		if firstLine {

			reqMethod, reqURI, err := parseRequestLine(line)
			if err != nil {
				log.Println(err)
				return
			}
			method, URI = reqMethod, reqURI

			firstLine = false
		}

		if line == "" {
			// when line is empty, http header is done
			log.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS (after that, body may also be read)")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("error while reading from connection: %v\n", err)
	}

	log.Println("finished parsing request message")

	bodyWriter := bufio.NewWriter(conn)

	// Basic Routing
	if method == "GET" && URI == "/" {
		rootPageTemplate.Execute(bodyWriter, nil)
	} else if method == "GET" && URI == "/apply" {
		applyPageTemplate.Execute(bodyWriter, nil)
	} else if method == "POST" && URI == "/apply" {
		fmt.Fprintln(bodyWriter, "<!DOCTYPE html>")
		fmt.Fprintln(bodyWriter, "<html>")
		fmt.Fprintln(bodyWriter, "<body>")
		fmt.Fprintln(bodyWriter, "  <header>")
		fmt.Fprintln(bodyWriter, "    <h1>POST APPLY</h1>")
		fmt.Fprintln(bodyWriter, "    <blockquote style=\"text-align: right\">\"HOLY COW THIS IS LOW LEVEL\"</blockquote>")
		fmt.Fprintln(bodyWriter, "  </header>")
		fmt.Fprintln(bodyWriter, "  <nav>")
		fmt.Fprintln(bodyWriter, "    <ul>")
		fmt.Fprintln(bodyWriter, "      <li><a href=\"/\">index</a></li>")
		fmt.Fprintln(bodyWriter, "      <li><a href=\"/apply\">apply</a></li>")
		fmt.Fprintln(bodyWriter, "    </ul>")
		fmt.Fprintln(bodyWriter, "  </nav>")
		fmt.Fprintln(bodyWriter, "</body>")
		fmt.Fprintln(bodyWriter, "</html>")
	} else {
		log.Println("no routes have been match")
	}

	sendOK(conn, bodyWriter)
}

// parseRequestLine parses the Request Line from the Http Request Message.
// Example value: "GET /foo/bar HTTP/1.1"
func parseRequestLine(line string) (method, URI string, err error) {
	fs := strings.Split(line, " ")
	if len(fs) != 3 {
		return "", "", fmt.Errorf("couldn't parse request line: '%v'", line)
	}

	method, URI, err = fs[0], fs[1], nil
	return
}

func sendOK(conn net.Conn, bodyWriter *bufio.Writer) {
	// https://tools.ietf.org/html/rfc7230#section-3.1.2
	//
	// HTTP Response Header
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")                        // HTTP Response Status Line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", bodyWriter.Buffered()) // HTTP Response Header Fields
	fmt.Fprint(conn, "Content-Type: text/html\r\n")

	// HTTP Response Whitespace Separator
	io.WriteString(conn, "\r\n")

	// HTTP Message Body
	bodyWriter.Flush()
}
