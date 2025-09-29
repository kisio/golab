package main

import "fmt"

// Function with no parameters
func greet() {
	fmt.Println("Hello!")
}

// Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple parameters of different types
func introduce(name string, age int) {
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}

func main() {
	// Call function with no parameters
	greet()

	// Call function with parameters and capture return value
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	// Call function with multiple parameters
	introduce("Alice", 25)
}
