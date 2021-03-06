package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// communicate through channel with values of type string
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// a value coming through channel c is a blocking call
	// this loop will block once a message comes and prints it out
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	/*
		could also be:

		for range links {
			fmt.Println(<-c)
		}
	*/
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up.")
	c <- "Yep it's up."
}
