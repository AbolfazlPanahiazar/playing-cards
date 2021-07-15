package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://github.com",
		"http://golang.org",
		"http://tenzumusic.com",
	}

	// create a channel
	c := make(chan string)

	// thread go routine
	for _, link := range links {
		go checkLink(link, c)
	}

	// recieve a vlue from the channel
	// for {
	// 	go checkLink(<-c, c)
	// }
	// or
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}
	fmt.Println(link + " is up!")
	c <- link
}
