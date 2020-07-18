package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://#",
		"https://#",
		"https://#",
		"https://#",
		"https://#",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	stat, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	} else if stat.StatusCode != 200 {
		fmt.Println(link, "has bad response")
		c <- link
		return
	} else {
		fmt.Println(link, "is up")
		c <- link
	}
	return
}
