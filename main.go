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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	//fmt.Println("Number of CPUs:", runtime.NumCPU()) //This will print the number of CPU cores available on your system
	//fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //This will print the current GOMAXPROCS value
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down, I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep, its up"
}
