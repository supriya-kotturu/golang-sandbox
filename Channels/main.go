package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links :=
		[]string{
			"https://google.com", "https://stackoverflow.com", "https://gnu.org",
		}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(27 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down =(")
		c <- link
	}
	fmt.Println(link, " is up!")
	c <- link
}
