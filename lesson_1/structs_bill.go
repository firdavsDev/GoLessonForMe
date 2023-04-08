package main

import (
	"os"
	"strconv"
)

/*
Map is a collection of key-value pairs.
Map vs Struct:
- Map is a collection of key-value pairs.
- Struct is a collection of fields.
- Map is a reference type.
- Struct is a value type.
- Map is not comparable.
- Struct is comparable.
- Map is not thread safe.
- Struct is thread safe.
- Map is a built-in type.
- Struct is a user-defined type.
- Map is a dynamic data structure.
- Struct is a static data structure.
- Map is a collection of unordered key-value pairs.
- Struct is a collection of ordered fields.

*/

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

	// b *Bill is a pointer to a Bill because we want to update the original Bill

}

// add_item adds an item to the bill
func (b Bill) add_item(name string, price float64) {
	b.items[name] = price

	// b Bill is a value of type Bill because we don't want to update the original Bill
}

// save saves the bill to a file
func (b *Bill) save() {
	b.update_is_paid(true)

	data := []byte(b.name)

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

}

func (b Bill) format() string {
	fs := "Bill breakdown: "

	for k, v := range b.items {
		fs += k + " - $" + strconv.FormatFloat(v, 'f', 2, 64) + " "
	}

	fs += "Total: $" + strconv.FormatFloat(b.total(), 'f', 2, 64)

	return fs
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
