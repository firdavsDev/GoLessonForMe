package main

import "fmt"

func variales() {

	// string
	var name string = "John"
	short_name := "John" // only works inside a function
	//update short_name := "John" // error
	short_name = "Jonibek" // works
	// short_name := "Jonibek" // error (short_name already declared)
	fmt.Println(name)
	fmt.Println(short_name)

	// int
	var age int = 30
	fmt.Println(age)

	// bool
	var isCool bool = true
	fmt.Println(isCool)

	// int shorthand
	// var age = 30
	age = 25
	fmt.Println(age)

	// bool shorthand
	// var isCool = true
	isCool = false
	fmt.Println(isCool)

	// bits & memory
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64 is unsigned (no negative numbers)
	// float32, float64 (decimal numbers) (float32 is less precise)
	// complex64, complex128 (complex numbers) (complex64 is less precise)

}

func main() {
	variales()
}
