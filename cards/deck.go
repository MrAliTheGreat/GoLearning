package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// We will just return the error of WriteFile - No handling
// 0666 --> Everyone can read and write
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	// These 2 lines below are for unique seed
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		// This will generate the same random value each time it's executed
		// newPosition := rand.Intn(len(d))
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
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
	For saving to file we can refer to the standard library of Go and use ioutil subpackage in io package
	There is a func named WriteFile. It will receive the data in a slice of bytes format
	So we need to convert our data to slice of bytes.
	If our data is of type string we can easily convert it to slice of bytes by casting it to []byte --> []bytes(data)
	In our case we have a deck type which is basically a slice of strings. We can use casting to simply convert deck back to []string
	So, we have to first make it a single string and then cast it to slice of bytes

	For making a single string out []string we can again refer to Go standard library and use strings package
	This package has a function named Join which like the join in Python
*/
// []byte is basically a string. []byte is the ascii value of each char of the string value. So the length of []byte is equal to number of chars in the string
/*
	For exiting the program we can use the os package of Go
	By call Exit and giving it the value of 0 the program will stop and the 0 will mean that everything went well
	But any non-zero value will show that there were some problems with the program, maybe an error occured
*/
/*
	For random numbers we can use the subpkg rand in pkg math
	Using the rand alone will result in the same random value each time since the seed is the same each time
	For creating a new seed each time we have to create a new rand instance and give a unique seed to it each time we use it for generating random values
	The code in shuffle func shows how to do so
	We can use the time for seed since each time we execute the program it's different
	Seed value must be of Int64 type so we have to use UnixNano to convert the value of time to Int64 type
*/
// VS Code must open on the working directory so that all the vars in different files can be recognized!
