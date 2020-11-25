package main

import "fmt"

func main() {
	/*
		WAP to find the average of total marks
	*/

	var marks [5]float32

	marks[0] = 98
	marks[1] = 89
	marks[2] = 85
	marks[3] = 70
	marks[4] = 81
	var total float32
	var i = 0
	for ; i < 5; i++ {
		total += marks[i]
	}
	fmt.Print("Total marks =", total)
	fmt.Println("Average marks =", total/float32(len(marks))) // TypeCasting
}
