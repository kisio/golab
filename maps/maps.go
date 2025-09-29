package main

import "fmt"

func main() {
	// Create a map using make
	m := make(map[string]int)

	// Add key-value pairs
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3

	// Access a value
	fmt.Println("Apple count:", m["apple"])

	// Check if a key exists
	if val, exists := m["orange"]; exists {
		fmt.Println("Orange count:", val)
	} else {
		fmt.Println("Orange not found in map")
	}

	// Iterate over the map
	fmt.Println("Map contents:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Delete a key
	delete(m, "banana")
	fmt.Println("After deleting banana:", m)

	// Length of map
	fmt.Println("Map length:", len(m))
}
