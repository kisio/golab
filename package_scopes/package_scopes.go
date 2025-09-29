package main

import "fmt"

// Package-level variable (global within package)
var globalVar = "I'm global"

// Package-level constant
const packageConst = 42

// Package-level function
func packageFunc() {
	fmt.Println("This is a package-level function")
}

func main() {
	// Accessing package-level variable
	fmt.Println("Global variable:", globalVar)

	// Accessing package-level constant
	fmt.Println("Package constant:", packageConst)

	// Calling package-level function
	packageFunc()
}
