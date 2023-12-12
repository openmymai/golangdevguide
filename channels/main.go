package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",		
	}

	c := make(chan string)

	for _, link := range links {
		// checkLink(link)
		go checkLink(link, c) // only go keyword with function, there is nothing
		// use channels to communicate between go routines
	}

	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<-c)
	// 	go checkLink(<-c)
	// }

	// for { // check forever
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// time.Sleep(5 * time.Second)
		// go checkLink(l, c)

		// Function literals
		go func(link string) { // add (link string) to prevent child routine to access the same address
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		// }()
		}(l) // Add l
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		// c <- "Might be down I think" // send message to channel c
		c <- link
		return 
	}

	fmt.Println(link, "is up!")
	// c <- "Yep it's up"
	c <- link
}