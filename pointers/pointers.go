package main

import "fmt"

func main() {
	// Declare a variable
	x := 10
	fmt.Println("x =", x)

	// Get address of x
	p := &x
	fmt.Println("Address of x (p) =", p)

	// Dereference pointer to get value
	fmt.Println("Value at p (*p) =", *p)

	// Modify value through pointer
	*p = 20
	fmt.Println("After *p = 20, x =", x)

	// Pointer to pointer
	pp := &p
	fmt.Println("Pointer to pointer (pp) =", pp)
	fmt.Println("Value at pp (**pp) =", **pp)

	// Nil pointer
	var nilPtr *int
	fmt.Println("Nil pointer:", nilPtr)
	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}
}
