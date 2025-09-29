package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Using bufio.Scanner for reading lines
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Hello,", name)

	// Using fmt.Scan for reading numbers
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("You are", age, "years old")

	// Reading multiple values
	var num1, num2 float64
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)
	fmt.Println("Sum:", num1+num2)
}
