package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

/*
	For holding multiple values Go has two data structures
	Array and slice
	Arrays and slices are basically the same but
	Arrays have fixed length
	Slices can grow or shrink in size and they have more features than arrays
	Elements in both of these data structures must be of same type
*/
// Append doesn't change the existing slice. It will create a new slice which then will be assigned to the exsiting slice
// Index for elements starts from 0 in Go
// Everytime we want to iterate through a slice or array we will use the keyword range
// Range will kinda make the slice loopable!
// In each iteration for the variables we will use :=
// The reason for using := is that in each iteration the previous value of the vars will be deleted and we need to declare new vars each iteration
// We can use _ for vars we won't use just like python!
/*
	For testing we have to name the test file with _test.go
	Meaning that if we want to test any code in a file named "a" we have to name the test file "a_test.go"
	Then after righting the test codes we can run all the tests in a package by running go test
*/
