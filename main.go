package main

import "fmt"

var x string = "Hello world" // golbally scoped

func main() {
	// Scope

	fmt.Println(x)
	p()
}

func p() {
	fmt.Println(x)
}
