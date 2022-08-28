package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// This way we have to rely on the order we put the attributes in the definition
	// Which is a bit risky
	alex := person{"Alex", "Anderson"}
	// This is certainly better!
	ali := person{firstName: "Ali", lastName: "Bahari"}
	fmt.Println(alex)
	fmt.Println(ali.firstName)
	fmt.Println(alex.lastName)

	var hossien person
	fmt.Println(hossien)
	hossien.firstName = "Hossein"
	hossien.lastName = "Moradi"
	fmt.Println(hossien)
}

/*
	In Go if we don't assign a value to a var it will give it a zero value by default
*/
