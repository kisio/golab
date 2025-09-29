package main

import "fmt"

// Function returning multiple values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Function with named return values
func calculate(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return // naked return
}

func main() {
	// Call function returning multiple values
	q, r := divide(10, 3)
	fmt.Printf("10 / 3 = %d, remainder %d\n", q, r)

	// Call function with named returns
	s, p := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)
}
