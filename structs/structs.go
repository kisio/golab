package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Define another struct with embedded struct
type Employee struct {
	Person // embedded struct
	ID     int
}

// Custom type (alias)
type MyInt int

func main() {
	// Create struct instance
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Person:", p.Name, p.Age)

	// Access fields
	p.Age = 31
	fmt.Println("Updated age:", p.Age)

	// Embedded struct
	e := Employee{
		Person: Person{Name: "Bob", Age: 25},
		ID:     123,
	}
	fmt.Println("Employee:", e.Name, e.Age, e.ID)

	// Custom type
	var num MyInt = 42
	fmt.Println("Custom type value:", num)
	fmt.Println("Underlying type:", int(num))
}
