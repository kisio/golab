package main

import "fmt"

func main() {
	var name string = "John Doe"
	var role string = "Software Engineer"
	fmt.Print("hello ")
	fmt.Print("human \n")
	fmt.Println("hello humans")
	fmt.Println("hello humans again", name)

	//formatted strings %_ is a format speciifier with _ being anything

	fmt.Printf("my name is  %q and i am a %q", name, role)
	fmt.Printf("role is of type %T \n", role)
	fmt.Printf("you scored %0.2f points \n", 100.00)

	//there are many format specifiers
	//Sprintf(save formatted strings)

	var str string = fmt.Sprintf("my name is  %q and i am a %q", name, role)
	fmt.Println("the saved string is ", str)
}
