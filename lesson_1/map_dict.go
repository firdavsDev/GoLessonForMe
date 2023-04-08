package main

import "fmt"

func main() {
	// create a map
	// the key is a string
	// the value is a string
	// the map is called m
	menu := map[string]float32{

		"pizza":  10.99,
		"burger": 5.99,
		"salad":  7.99,
	}

	for k, v := range menu {
		fmt.Printf("%s is key: $%.2f is value \n", k, v)
	}
}
