/*
Shorter hand syntax for creating an array
*/

package main

import "fmt"

func main() {
	x := [5]float64{
		98,
		99,
		100,
		121,
		45,
	}

	for _, value := range x {
		fmt.Print(value, " ")
	}
}
