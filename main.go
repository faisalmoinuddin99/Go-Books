package main

import "fmt"

func main() {
	var x string
	x = "first value"
	fmt.Println(x)

	x = x + " second value"
	fmt.Println(x)

	// OR using special Assignments

	x += " Third Value"
	fmt.Println(x)
}
