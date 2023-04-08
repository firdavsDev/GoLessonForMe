package main

import (
	"os"
	"strconv"
)

type Bill struct {
	name    string             // name of the bill
	items   map[string]float64 // items in the bill
	is_paid bool
}

// total calculates the total amount of the bill
func (b Bill) total() float64 { // it is method of Bill type
	total := 0.0
	for _, v := range b.items {
		total += v
	}
	return total
}

// update_is_paid updates the status of the bill
func (b *Bill) update_is_paid(status bool) {
	b.is_paid = status

}

// add_item adds an item to the bill
func (b Bill) add_item(name string, price float64) {
	b.items[name] = price

}

// Formater
func (b Bill) format() string {
	fs := "Bill breakdown: \n\n"

	for k, v := range b.items {
		fs += k + " - ...$" + strconv.FormatFloat(v, 'f', 2, 64) + "\n"
	}
	fs += "----------------------------------\n"
	fs += "Total: $" + strconv.FormatFloat(b.total(), 'f', 2, 64)

	return fs
}

// save saves the bill to a file
func (b *Bill) save() string {

	b.update_is_paid(true) // update the status of the bill

	data := []byte(b.format()) // convert the name of the bill to a byte slice

	file_path := "saved_bills/" + b.name + ".txt"

	err := os.WriteFile(file_path, data, 0644)

	if err != nil {
		panic(err)
	}

	return file_path
}

// NewBill creates a new bill with the given name
func NewBill(name string) Bill {
	b := Bill{
		name:    name,
		items:   map[string]float64{},
		is_paid: false,
	}
	return b
}
