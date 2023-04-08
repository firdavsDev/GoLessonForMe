package main

import "fmt"

func main() {
	// for loop (traditional)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// // while loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

	names := []string{"John", "Paul", "George", "Ringo"}
	for index, name := range names {
		fmt.Println(index, name)
	}

}
