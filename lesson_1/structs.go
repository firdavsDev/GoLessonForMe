package main

import "fmt"

func main() {
	my_bill := NewBill("my bill") // my_bill is a value of type Bill

	fmt.Println(my_bill)
	fmt.Println(my_bill.total())
}
