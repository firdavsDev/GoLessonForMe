package main

import "fmt" // Format package (fmt) is used to print to the terminal

func printer() {
	fmt.Println("Hello, World!")
	name := "Jonibek"
	age := 30
	salary := 100.5

	fmt.Println("My name is", name, "and I am", age, "years old.")
	fmt.Printf("My name is %v and I am %v years old. \n", name, age)
	fmt.Printf("Age type: %T \n", age)
	fmt.Printf("Salary type: %0.2f \n", salary) // 0.2f means 2 decimal places after the dot

	// save SprintF to a variable
	var_str := fmt.Sprintf("My name is %v and I am %v years old. \n", name, age)
	fmt.Println(var_str)
}

func main() {
	printer()
}
