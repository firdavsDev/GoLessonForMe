package main

import (
	"fmt"
	"sort"
)

func main() {
	// strings
	// greetings := "Hello, World!"
	// fmt.Println(greetings)
	// fmt.Println(strings.Contains(greetings, "Hello"))         // true (case sensitive)
	// fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi")) // Hi, World!
	// fmt.Println(strings.ToUpper(greetings))                   // HELLO, WORLD!
	// fmt.Println(strings.ToLower(greetings))                   // hello, world!
	// fmt.Println(strings.ToTitle(greetings))                   // HELLO, WORLD!
	// fmt.Println(strings.Index(greetings, "World"))            // 7
	// fmt.Println(strings.Split(greetings, ","))                // [Hello  World!]
	// fmt.Println(strings.Trim(greetings, "Hello, "))           // World!

	// sort int
	// ages := []int{4, 2, 6, 1, 3, 7, 9, 5, 8, 10}
	// sort.Ints(ages)   // sort in ascending order
	// fmt.Println(ages) // [1 2 3 4 5 6 7 8 9 10]
	// index := sort.SearchInts(ages, 5) // search for the index of the value 5 in the sorted array ages if it exists it will return the index of the value 5
	// fmt.Println(index) // 5

	// sort string
	names := []string{"Jonibek", "Alisher", "Samandar", "Sardor"}
	sort.Strings(names) // sort in ascending order
	fmt.Println(names)  // [Alisher Jonibek Samandar Sardor]

}
