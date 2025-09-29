package main

import "fmt"

// Define a struct
type Rectangle struct {
	Width, Height int
}

// Method with value receiver
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Method with value receiver that doesn't modify
func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Rectangle:", rect)
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}
