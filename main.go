// Package main is the entry point for the application.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"https://microsoft.com",
		"https://openai.com/",
		"https://go.dev/",
		"https://www.php.net/",
	}

	for _, link := range links {
		//fmt.Println(link)
		checkLink(link)
	}
}

// silly function that does not care about the response,
// but rather focuses on whether there is an error or not
// when performing an HTTP request
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
