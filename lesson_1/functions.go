package main

import "fmt"

func calculateBill(price, no int) int {
	totalPrice := price * no
	return totalPrice
}

func greeting(name string) string {
	return "Hello " + name
}

func goodbye(name string) string {
	return "Goodbye " + name
}

func cycle_names(names []string, name_f_function func(string) string) []string {
	returend_values := []string{}
	for _, name := range names {
		returend_values = append(returend_values, name_f_function(name))
	}
	return returend_values
}

func main() {
	// function call
	// totalPrice := calculateBill(10, 5)
	// fmt.Println(totalPrice)

	names := []string{"John", "Paul", "George", "Ringo"}
	list_of_greetings := cycle_names(names, greeting)
	fmt.Println(list_of_greetings)
	list_of_goodbyes := cycle_names(names, goodbye)
	fmt.Println(list_of_goodbyes)
	fmt.Println(len(list_of_goodbyes))
	fmt.Printf("%T \n", list_of_goodbyes)

}
