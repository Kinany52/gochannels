// Package main is the entry point for the application.
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"https://microsoft.com",
		"https://openai.com/",
		"https://go.dev/",
		"https://www.php.net/",
	}

	c := make(chan string)
	var wg sync.WaitGroup

	// Start goroutines to check each link
	for _, link := range links {
		wg.Add(1) // Increment the counter for each goroutine
		go func(link string) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			checkLink(link, c)
		}(link)
	}

	// Start a goroutine to close the channel after all links are processed
	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(c)  // Close the channel after all goroutines have finished
	}()

	// Process results
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

	//fmt.Println("Number of CPUs:", runtime.NumCPU())
	//fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
