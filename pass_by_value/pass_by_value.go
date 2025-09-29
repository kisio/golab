package main

import "fmt"

// Function that tries to modify an int (passed by value)
func modifyInt(x int) {
	x = 10
	fmt.Println("Inside modifyInt, x =", x)
}

// Function that modifies a slice (reference to underlying array)
func modifySlice(s []int) {
	s[0] = 99
	fmt.Println("Inside modifySlice, s =", s)
}

func main() {
	// Pass by value for int
	num := 5
	fmt.Println("Before modifyInt, num =", num)
	modifyInt(num)
	fmt.Println("After modifyInt, num =", num) // unchanged

	// Pass by value for slice (but slice contains pointer to array)
	slice := []int{1, 2, 3}
	fmt.Println("Before modifySlice, slice =", slice)
	modifySlice(slice)
	fmt.Println("After modifySlice, slice =", slice) // changed
}
