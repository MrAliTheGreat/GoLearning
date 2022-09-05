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

	// Creating a channel of type string. So only string can be passed through the channel between go routines
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Another way of saying the infinite for loop from previous commit
	for l := range c {
		go checkLink(l, c)
	}
	/*
		range c means that wait for the channel to return sime value
		Then after the channel has returned some value assign it to var l
	*/
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

	// When link goes through the channel to the main go routine the main will understand which link is finished so it will call the func with that link again in a new go routine
}

/*
	The sequential habit of this program is a problem!
	If the number of sites is in the thousands then the delay for each request can add up to a large amount --> slower response
	We have to wait until we get to a certain site since the previous ones have to finish --> Checking status in different times instead of an exact time
	So, let's go parallel!
*/
/*
	Our whole code is one go routine that executes each line of code one by one --> The main go routine is created when we launch our program
	When we use a blocking call the main go routine stops and can do nothing else
	So, for fixing this we can launch a new go routine!
	We will use go when calling a func and this will run the func in a new go routine
	When a go routine is calling a blocking statement everyone is informed so that can run in the meantime. Maybe a new go routine will be created or maybe another existing go routine will continue

	Go has a Go scheduler. Go Scheduler in one core cpus will run only one go routine at a time
	When a go routine is finished or running a blocking code the scheduler will tell the go routine you have had enough and will give the one core cpu to another go routine of its choice
	So, truly the go routines are not running at the same time but since the context switch is so fast it seems like every go routine is executing at the same time --> Concurrency
	Our program is CONCURRENT if it has the ability to load up multiple go routines at a time

	Go in default tries to use only one core of cpu! for using more than one cpu we have to overwrite some settings!
	Each core can run a go routine and since we have multiple of them the go scheduler will assign each a go routine to each core
	This time each go routine is truly running at the same time --> Parallelism
	We only get PARALLELISM once we start to include multiple physical cores on our machine
	Run Go routine at the very exact time (requires multiple cores on the cpu, one go routine on one core at the same time) --> PARALLELISM

	Go scheduler will work for each core meaning that for each core the go scheduler will decide which go routine to choose and execute
	All other go routines that we create with the keyword "go" will be child go routines! (Instead of main go routine that is created once we run our program)
	The main go routine is the one that controls when our program ends. When there is nothing more to run, the main go routine will simply stop the program
	The main go routine won't care if there is any other go routine (child go routines) still running when it reaches the end!

	The keyword "go" can only be used before func calls
*/
/*
	Channels are used to communicate in between different go routines
	Channel is the only way of communication between go routines
	The data we attempt to share via a channel must all be of the same type
	One channel can be used to make communication between main routine and other child routines
	You can think of a channel as a group chat. Every go routine can send the same type of data to any other go routine via channels
	channel <- data : Sending data into the channel
	varName <- channel : Receiving data from channel and putting it in var varName
	Really important to know: Receiving a data through a channel ( <-channel ) is a BLOCKING call!
*/
