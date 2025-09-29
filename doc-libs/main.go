package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	greeting := "hello there human"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println("Square root of 16:", math.Sqrt(16))
	fmt.Println("Current time:", time.Now())
	fmt.Println("Home directory:", os.Getenv("HOME"))

	// Using bufio
	reader := bufio.NewReader(strings.NewReader("Hello from bufio"))
	line, _ := reader.ReadString('\n')
	fmt.Println("Bufio read:", strings.TrimSpace(line))

	// Using encoding/json
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{Name: "Alice", Age: 25}
	jsonData, _ := json.Marshal(person)
	fmt.Println("JSON:", string(jsonData))

	// Using io
	src := strings.NewReader("Hello from io")
	dst := &strings.Builder{}
	io.Copy(dst, src)
	fmt.Println("IO copy result:", dst.String())

	// Using log
	log.Println("This is a log message")

	// Using net/http
	resp, err := http.Get("https://httpbin.org/get")
	if err == nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("HTTP response length:", len(body))
	} else {
		fmt.Println("HTTP request failed")
	}
}
