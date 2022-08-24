package main

import "fmt"

// This is like inheritance. Deck has all the functionality of a slice of strings
type deck []string

/*
	This is a function with a receiver!
	What it means is that any var of type deck gets access to print method

	(d deck) is the receiver part anything other than that is exactly like a normal func
	deck is the type that we want to attach the print function to
	d is like the instance of deck that we are working with like the actual copy
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
		fmt.Println("=====")
	}
}

// This deck.go file is like when we want to create (declare) a class in a seperate file
// Everything will be run through the main.go file
/*
	Go is not a OO language so we have to extend what we have to kinda work like OO
	For attributes we will use type or struct
	For methods we will use function with a receiver
	So we will create a new type which is created from base types then we will create custom methods only for this new type with function with receiver
*/
/*
	If there are multiple files in the package, while running the program all the file names must be provided
	So something like --> go run *.go --> THIS DOESN'T WORK! Must provide all the names individually
*/

// VS Code must open on the working directory so that all the vars in different files can be recognized!
