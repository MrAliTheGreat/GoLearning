package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

// We can remove contact and just leave contactInfo. Go will automatically create a field with the name contactInfo
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Johnson",
		contact: contactInfo{
			email: "Alex@yahoo.com",
			zip:   123,
		},
	}

	alex.print()
}

func (p person) print() {
	fmt.Println("First Name: " + p.firstName)
	fmt.Println("Last Name: " + p.lastName)
	fmt.Println("Email: " + p.contact.email)
	fmt.Printf("zip: %d", p.contact.zip)
}

// In Go if we don't assign a value to a var it will give it a zero value by default
// When initializing a struct, all lines must have a , at the end of them. Even the last value
// "%+v" in printf will print all the attribute names and their values in console
