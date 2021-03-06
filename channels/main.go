package main

import (
	"fmt"
	"net/http"
	"time"
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

	// wait for the channel to return a value and then assign it to l
	// i.e. call checkLink for each link forever
	for l := range c {
		// use function literal to pause only when this infinite loop does an iteration,
		// not during the initial loop above
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up.")
	c <- link
}
