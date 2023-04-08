package main

import "fmt"

// go run structs.go structs_bill.go

func main() {

	mybill := createBill() //

	promtOptions(mybill) //inputer/get_data.go

	fmt.Println(mybill)

}
