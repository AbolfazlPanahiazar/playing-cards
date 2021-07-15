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
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		c <- "Might be down i think"
		return
	}
	fmt.Println(link + " is up!")
	c <- "Yep its up"
}
