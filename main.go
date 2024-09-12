// Package main is the entry point for the application.
package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU()) //This will print the number of CPU cores available on your system.
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0)) //This will print the current GOMAXPROCS value, which is the number of CPU cores Go is currently using.
	//The value won't be set to 0; it will simply return the number of OS threads Go is currently using (which, by default, is set to the number of available CPU cores). If you want to change it, you need to pass a non-zero value to GOMAXPROCS.
	return

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
