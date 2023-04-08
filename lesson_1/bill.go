package main

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

func (b Bill) total() float64 { // it is method of Bill type
	total := 0.0
	for _, v := range b.items {
		total += v
	}
	return total
}

// NewBill creates a new bill with the given name
func NewBill(name string) Bill {
	b := Bill{
		name: name,
		items: map[string]float64{
			"pie":       5.99,
			"ice cream": 3.99,
		},
		is_paid: false,
	}
	return b
}
