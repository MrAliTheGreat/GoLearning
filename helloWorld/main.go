// Compiling and Running the program --> go run main.go
// Compiling the program --> go build main.go --> This will create an executable file
// Formating the code --> go fmt
// Installing a package --> go install
// Downloading raw source code of a package --> go get
// Running tests --> go test

package main

import "fmt"

func main() {
	fmt.Println("Well Hello There!")
}

// Package = Project
// A package can have multiple .go files
/*
	Packages can be executable or reusable
	Executable package when compiled will created an executable file
	Reusable packages are like libraries or helper function
	The word main will make a package executable meaning a package named main will be executable
	If a package is named anything other than main it will be considered reusable and when compiled will not create an executable file
	An executable package must have a main func, it is the place where the code execution will start
*/
