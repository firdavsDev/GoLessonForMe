package main

import "fmt"

func main() {
	age := 18

	// fmt.Println(age == 18)
	// fmt.Println(age != 18)
	// fmt.Println(age > 18)
	// fmt.Println(age < 18)
	// fmt.Println(age >= 18)
	// fmt.Println(age <= 18)

	// if statement
	if age >= 18 && age <= 60 {
		fmt.Println("You are an adult")
	} else if age > 60 {
		fmt.Println("You are a senior")
	} else {
		fmt.Println("You are a child")
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, number := range numbers {

		if index == 1 {
			fmt.Printf("%v index is %v continue her \n", number, index)
			continue
		} else if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}

	}

}
