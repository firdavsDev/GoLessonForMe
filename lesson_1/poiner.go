package main

/*

Pointer is a variable that stores the memory address of another variable.

& is address operator and * is dereference operator

####################################################################
# memory address   # value        # pointer #   variable 		 #
# 0x1234           # 42           # 0x1234  #   a = 42			 #
# 0x1234           # 42           # 0x1234  #   b = &a			 #
# 0x1234           # 43           # 0x1234  #   *b = 43			 #
####################################################################


Non pointer value is copied when passed to a function.
There are: string int float bool array struct

Pointer Wrapper: slice map channel Function , they are not copied when passed to a function.

*/
import "fmt"

func main() {
	var a int = 42
	b := &a // b is pointer to a memory address. & is address operator and * is dereference operator

	fmt.Printf("%v is value and %v is memory addrees\n", a, b)

	// Use * to read val from address
	fmt.Println(*b) // 42
	fmt.Println(&b)

	// Change val with pointer
	*b = 43 // change value of a to 43
	fmt.Println(a)
}
