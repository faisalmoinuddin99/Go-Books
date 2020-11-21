/*
Here’s an example program that takes in a number entered by the user and
doubles it:
*/

package main

import "fmt"

func main() {
	fmt.Print("Enter the number:")
	var userInput float64
	fmt.Scanf("%f", &userInput)

	userInput *= 2
	fmt.Println("The double of the entered value is:", userInput)
}

/*
Output : go run main.go
Enter the number:4
The double of the entered value is: 8
*/
