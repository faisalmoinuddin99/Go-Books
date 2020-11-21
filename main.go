package main

import "fmt"

func main() {
	// Equality Operator

	var x string = "Hello"
	var y string = "world"

	fmt.Println(x == y)

	var z string = "world"

	fmt.Println(y == z)
}
