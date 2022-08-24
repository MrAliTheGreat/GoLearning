package main

import "fmt"

func newString() string {
	return "Card 4"
}

func main() {
	// Different ways to declare a variable
	var card_1 string = "Card 1"
	var card_2 = "Card 2"
	card_3 := "Card 3"
	card_4 := newString()

	fmt.Println(card_1)
	fmt.Println(card_2)
	fmt.Println(card_3)
	fmt.Println(card_4)
}
