package main

import "fmt"

// Define a struct
type Counter struct {
	Value int
}

// Method with pointer receiver (can modify the struct)
func (c *Counter) Increment() {
	c.Value++
}

// Method with pointer receiver for setting value
func (c *Counter) SetValue(newValue int) {
	c.Value = newValue
}

func main() {
	c := Counter{Value: 0}
	fmt.Println("Initial counter:", c.Value)

	// Call method with pointer receiver
	c.Increment()
	fmt.Println("After increment:", c.Value)

	c.SetValue(10)
	fmt.Println("After set value:", c.Value)
}
