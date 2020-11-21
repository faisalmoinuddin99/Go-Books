/*
write a program that converts from Fahrenheit into Celsius
(C = (F − 32) * 5/9).
*/

package main

import "fmt"

func main() {
	/*
		Todos:
		1. Take input as in Fahrenheit
		2. convert it into celsius
		3. display the output in celsius
	*/

	var Fahrenheit float64
	var Celsius float64

	fmt.Print("Enter the Temprature in Fahrenheit :")
	fmt.Scanf("%f", &Fahrenheit)
	Celsius = ((Fahrenheit - 32) * 5 / 9)
	fmt.Print("The conversion result = ", Celsius, "°C")
}

/*
Output :go run main.go
Enter the Temprature in Fahrenheit :100
The conversion result = 37.77777777777778°C
*/
