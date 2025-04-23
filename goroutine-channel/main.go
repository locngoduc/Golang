package main

import (
	"fmt"
	"net/http"
)

func fetch(url string, ch chan string) {
	resp, _ := http.Get(url)
	fmt.Println("Fetching:", url)
	ch <- fmt.Sprintf("Fetched %s: %d", url, resp.StatusCode)
}

func main() {
	ch := make(chan string)
	urls := []string{"https://example.com", "https://golang.org", "https://google.com"}

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		result := <-ch
		fmt.Println(result)
	}
}
