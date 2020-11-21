/*
Write another program that converts from feet into meters
(1 ft = 0.3048 m).
OR
(divide value of foot with 3.281)
*/

package main

import "fmt"

func main() {
	/*
		Todos:
		1. Take input as in feet
		2. convert it into meters
		3. display the output in meters
	*/

	var feet float64
	var meters float64

	fmt.Print("Enter the value in Feet :")
	fmt.Scanf("%f", &feet)
	meters = (feet / 3.281)
	fmt.Print("The conversion result = ", meters, "m")
}

/*
Output :go run main.go
Enter the value in Feet :1
The conversion result = 0.30478512648582745m
*/
