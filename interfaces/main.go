package main

import "fmt"

// This means that any type that has a func of getGreeting() which returns a string is also a bot
// We can also provide the arguments that will be the input of the func in the ()
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// CUSTOM logic for each language
// If we do not use the receiver var we can just omit it and leave the type
func (englishBot) getGreeting() string {
	return "Well hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola Hector!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// In Go funcs with the same name but different arguments are not permitted
