package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://bing.com",
		"http://amazon.com",
		"http://github.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!")
		return
	}

	fmt.Println(link + " is up!")
}

/*
	The sequential habit of this program is a problem!
	If the number of sites is in the thousands then the delay for each request can add up to a large amount --> slower response
	We have to wait until we get to a certain site since the previous ones have to finish --> Checking status in different times instead of an exact time
	So, let's go parallel!
*/
