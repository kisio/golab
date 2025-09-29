package main

import "fmt"

func main() {

	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	// Range loop over slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Infinite loop with break
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println("Infinite loop count:", count)
		count++
	}
}
