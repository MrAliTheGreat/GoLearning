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

	alex.updateFirstName("ali")
	alex.print()
	fmt.Println("===== Reference Type =====")

	test := []string{"a", "b", "c", "d"}
	changeSlice(test)
	for _, val := range test {
		fmt.Print(val + ", ")
	}
}

func changeSlice(slice []string) {
	// No pointers but the change will happen since slice is a reference type (pointer underneath)
	slice[0] = "z"
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Println("First Name: " + p.firstName)
	fmt.Println("Last Name: " + p.lastName)
	fmt.Println("Email: " + p.contact.email)
	fmt.Printf("zip: %d\n", p.contact.zip)
}

// In Go if we don't assign a value to a var it will give it a zero value by default
// When initializing a struct, all lines must have a , at the end of them. Even the last value
// "%+v" in printf will print all the attribute names and their values in console
/*
	Go is a pass by value language. When a value is passed to a func, Go copies that value and the copy, NOT the main value, is available for use in the func.
	So, if we want to change the value we have to pass by reference --> Pointer baby! Pointers!
*/
/*
	When defining a function with a pointer receiver, Go will allow access to both the main type and its pointer
	Meaning that if we define the receiver as a pointer, we can call the func using the main obj or its pointer
	Like if the receiver is like: func (a *obj) test(){} --> we can use this func by both: objInstance.test and &objInstance.test
	But since defining a new var each time an assign it &objInstance is dumb, one can simply use objInstance and have access to the func with pointer receiver
	So the main type (obj) has access to funcs with normal (a obj) and pointer (a *obj) receiver
	Although there will be a mismatch, Go allows this to happen. Go will automatically create the required pointer. Great shortcut ngl!
*/
/*
	Slice is a data structure that consists of 1) Pointer to head of array 2) Capacity of array (how many elements can it contain) 3) Length of the array (current num of exisiting elements)
	So, slice is kind of an array to be honest.

	With this in mind that a slice is basically a pointer to an array, we can understand why slices can be changed without passing them via their pointers
	So Go does its job correctly. When a slice is passed to a func it will make copy of it because of the whole pass by value thing
	But the copy of the slice still is a pointer and it points to the same location as the main slice, the array. So, when we change the values in the slice the change will happen on the main array

	Important thing to note here is that slices are of reference type
	Meaning that they just refer to another data in the memory. Like we see here that slice points to an array and by itself it is just the ds I explained above
	So in total, when using slices and reference types in general we don't need to worry about passing them as a pointer. But in value types like structs, int, string ..., we have to be carful about this whole pass by value thing
*/
