package main

import (
	"fmt"
	"strings"
)

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
	for _, card := range d {
		fmt.Println(card)
	}
}

/*
We don't need a receiver here since we are creating a new deck
Like there are not any deck so we can't give access to nothing! we have to create it first then give access to it via receiver

In deck{} we have to put {} to make an empty slice. Without it the slice will NOT be initialized to empty slice by default
*/
func newDeck() deck {
	cards := deck{}

	numbers := []string{"Ace", "Two", "Three", "Four"}
	marks := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, num := range numbers {
		for _, mark := range marks {
			cards = append(cards, num+" of "+mark)
		}
	}

	return cards
}

// Go has support for returning multiple values in a function like (deck, deck) here
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// This deck.go file is like when we want to create (declare) a class in a seperate file
// Everything will be run through the main.go file
// For importing multiple pkgs we can use () and mention each pkg in it without any comma. It's better than writing import each time!
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
/*
	For saving to file we can refer to the standard library of Go and use ioutil package
	There is a func named WriteFile. It will receive the data in a slice of bytes format
	So we need to convert our data to slice of bytes.
	If our data is of type string we can easily convert it to slice of bytes by casting it to []byte --> []bytes(data)
	In our case we have a deck type which is basically a slice of strings. We can use casting to simply convert deck back to []string
	So, we have to first make it a single string and then cast it to slice of bytes

	For making a single string out []string we can again refer to Go standard library and use strings package
	This package has a function named Join which like the join in Python
*/

// VS Code must open on the working directory so that all the vars in different files can be recognized!
