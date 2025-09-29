// arrays and slices
package main

import "fmt"

func main() {
	// //arrays are of fixed length and cant change
	// var ages[5]int = [5] int {23,34,56,78,99}
	var ages = [3]int{23, 45, 56}
	fmt.Println(ages, len(ages))

	//shorthand
	names := [3]string{"John", "Jane", "Alice"}
	fmt.Println(names, len(names))

	//slices
	var scores = []int{100, 34, 546}
	scores[2] = 344
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[0:]
	rangeTwo := names[1:]
	fmt.Println(rangeOne, rangeTwo)

	rangeOne = append(rangeOne, "Bob")
	fmt.Println(rangeOne)
}
