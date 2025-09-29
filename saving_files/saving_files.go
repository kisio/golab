package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Using os.WriteFile for simple writing
	data := "Hello, World!\nThis is a test file."
	err := os.WriteFile("output.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File saved successfully using os.WriteFile")

	// Using bufio.Writer for more control
	file, err := os.Create("output2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Line 1\n")
	writer.WriteString("Line 2\n")
	writer.Flush() // Important to flush
	fmt.Println("File saved successfully using bufio.Writer")
}
