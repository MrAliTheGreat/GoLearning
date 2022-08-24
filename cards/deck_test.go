package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length 16, but got %v", len(d))
		// t.Errorf("Expected length 16, but got " + fmt.Sprint(len(d)))
	}
	// Other tests can be written here but I think one is enough :)
}

/*
	In testing if something goes wrong we'll tell it to the test handler
	*testing.T is the test handler and is required by go test
	go test didn't work at first. It would show a message about not being able to find main module
	I solved this by setting the Go variable of GO111MODULE to auto using the command below
		go env -w GO111MODULE=auto
*/
/*
	os.Remove will delete a file or directory if it doesn't exist it will return an error
	If we just want to delete the file and the file doesn't exist then not handling the returned error will make everything go without a problem
	Meaning the error will say the file isn't there and you would say OK that's what I want my bad for calling Remove.
	The important thing is that the program by default won't halt just because the file is not there
*/
