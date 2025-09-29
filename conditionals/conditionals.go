package main

import "fmt"

func main() {

	// Boolean variables
	isTrue := true
	isFalse := false
	fmt.Println("isTrue:", isTrue, "isFalse:", isFalse)

	// If statement
	if isTrue {
		fmt.Println("This is true")
	}

	// If-else
	age := 20
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// If-else if-else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// Switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}
}
