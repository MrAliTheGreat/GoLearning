package main

import "fmt"

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

// SIMILAR logic in both languages
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// In Go funcs with the same name but different arguments are not permitted
