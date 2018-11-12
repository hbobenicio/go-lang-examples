package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	links := []string{
		"http://amazon.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	checkLinks(links)
	checkLinks2(links)

	fmt.Println("----")
	fmt.Println("done scanning links")
}

func checkLinks(links []string) {
	var wg sync.WaitGroup
	wg.Add(len(links))

	for _, link := range links {
		go func(link string) {
			defer wg.Done()

			checkLink(link)
		}(link)
	}

	wg.Wait()
}

func checkLink(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s might be down: %v\n", link, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, resp.Status, link)
		return
	}

	fmt.Println(resp.Status, link)
}

// Example 2
type checkLinkResponse struct {
	Link string
	Err  error
}

func checkLinks2(links []string) {
	var responses []chan checkLinkResponse
	for _, link := range links {
		responses = append(responses, goCheckLink2(link))
	}

	var wg sync.WaitGroup
	wg.Add(len(responses))
	for i := 0; i < len(responses); i++ {
		go func(chIndex int) {
			defer wg.Done()

			resp := <-responses[chIndex]
			if resp.Err != nil {
				fmt.Fprintln(os.Stderr, resp.Err)
			} else {
				fmt.Println(resp.Link, "is up")
			}
		}(i)
	}
	wg.Wait()
}

func goCheckLink2(link string) chan checkLinkResponse {
	resp := make(chan checkLinkResponse)
	go func() {
		defer close(resp)
		resp <- checkLink2(link)
	}()
	return resp
}

func checkLink2(link string) checkLinkResponse {
	resp, err := http.Get(link)
	if err != nil {
		e := fmt.Errorf("%s might be down: %v", link, err)
		return checkLinkResponse{Link: link, Err: e}
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, resp.Status, link)
		e := fmt.Errorf("%s %d", link, resp.StatusCode)
		return checkLinkResponse{Link: link, Err: e}
	}

	return checkLinkResponse{Link: link}
}
