package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://youtube.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// variant 1
	// for index := 0; index < len(links); index++ {
	// 	fmt.Println(<-c)
	// }

	// variant 2
	// for {
	// 	go checkLink(<-c, c)
	// }

	// variant 3
	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	// variant 4
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
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
	fmt.Println(link, "is up!")
	c <- link
}
