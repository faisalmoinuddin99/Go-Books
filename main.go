package main

import "fmt"

func main() {
	/*
		Switch Case:
		WAP to that printed the English names for numbers.
	*/

	var i int16
	i = 1
	for i <= 6 {
		switch i {
		case 0:
			fmt.Print("Zero")
			break
		case 1:
			fmt.Print("One")
			break
		case 2:
			fmt.Print("Two")
			break
		case 3:
			fmt.Print("Three")
			break
		case 4:
			fmt.Print("Four")
			break
		case 5:
			fmt.Print("Five")
			break
		case 6:
			fmt.Print("Six")
			break
		default:
			fmt.Print("unknown number")
			break

		}

		i++
	}
}
