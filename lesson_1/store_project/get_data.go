package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promtOptions(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := get_input("Choose option (a - add item, s - save bill): ", reader)

	switch opt {
	case "a":
		name, _ := get_input("Item name: ", reader)
		price, _ := get_input("Item price: ", reader)
		// load the price into a float64
		price_float, err := strconv.ParseFloat(price, 64) // 64 is the number of bits
		if err != nil {
			fmt.Println("the price must be a number")
			promtOptions(b)
		}
		b.add_item(name, price_float)
		fmt.Println("item added -", name, price)
		promtOptions(b)
	case "s":
		file_path := b.save()
		fmt.Println("you saved the bill in ", file_path)
	default:
		fmt.Println("that was not a valid option")
		promtOptions(b)
	}
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	billName, _ := get_input("Create a new bill name: ", reader)

	b := NewBill(billName) // ./model/bill.go

	return b
}
