package main

import "fmt"

func main() {
	// The first string is for keys type and the seconds string is for values type
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}

	colors["white"] = "#fffff"

	delete(colors, "black")

	fmt.Println(colors)
	fmt.Println("=============")
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key + ": " + value)
	}
}

// In maps all the keys are of the same type. This is also the case fo values
/*
	In maps all keys must be of the same type. Also all values must be of same type
	In struct values can be of different type

	In maps keys are indexed. So, we can iterate over them
	In struct keys don't support indexing. So, no iteration!

	Map is a reference type
	Struct is a value type

	We should use maps when we want to represent a collection of related properties
	We should use structs when we want to represent a "thing" with a lot of different properties

	In maps we don't need to know all the keys at compile time. We can add or remove key value pairs over time
	In structs we have to know all the properties at the compile time and they can not change over time
*/
