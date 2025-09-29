package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Parse string to float64
	str := "3.14159"
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
	} else {
		fmt.Println("Parsed float64:", f)
	}

	// Parse with different bit sizes
	str2 := "2.718"
	f32, err := strconv.ParseFloat(str2, 32)
	if err != nil {
		fmt.Println("Error parsing float32:", err)
	} else {
		fmt.Println("Parsed float32:", float32(f32))
	}

	// Handling invalid input
	invalid := "not a number"
	_, err = strconv.ParseFloat(invalid, 64)
	if err != nil {
		fmt.Println("Expected error for invalid input:", err)
	}
}
