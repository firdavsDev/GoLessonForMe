package main

import "fmt"

func main() {
	// Define a string array with a length of 5
	// var fruits [5]string
	// fruits[0] = "Apple"
	// fruits[1] = "Orange"
	// fruits[2] = "Grape"
	// fruits[3] = "Banana"
	// fruits[4] = "Pineapple"
	// fmt.Println(fruits)
	// fmt.Println(fruits[1])

	// Declare and assign values
	fruits := [5]string{"Apple", "Orange", "Grape", "Banana", "Pineapple"}
	fmt.Println(fruits)      // [Apple Orange Grape Banana Pineapple]
	fmt.Println(fruits[1])   // index starts from 0
	fmt.Println(len(fruits)) // length of the array

	// Slice
	fruits_slice := []string{"Apple", "Orange", "Grape", "Banana", "Pineapple"}
	fmt.Println(fruits_slice)      // [Apple Orange Grape Banana Pineapple]
	fmt.Println(fruits_slice[1:3]) // index starts from 0 (1:3 means 1 to 3, but not including 3)
	fmt.Println(len(fruits_slice)) // length of the array

	// Update slice
	fruits_slice[0] = "Strawberry"
	fmt.Println(fruits_slice)                    // [Strawberry Orange Grape Banana Pineapple]
	fruits_slice = append(fruits_slice, "Mango") // append to the end of the slice (Mango)
	fmt.Println(fruits_slice)                    // [Strawberry Orange Grape Banana Pineapple Mango]

	// Slice range
	fruits_slice = append(fruits_slice[1:3], fruits_slice[4:]...) // append to the end of the slice (Mango)
}
