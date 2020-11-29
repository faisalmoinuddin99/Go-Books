/*
Slice
*/

package main

import "fmt"

func main() {

	//Write a program that finds the smallest number in this list:

	list := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	//var sizeOfList = len(list)

	var min = 0
	min = list[0]
	fmt.Println("Value of min before loop :", min)
	for _, value := range list {
		if value < min {
			min = value
		}

	}
	fmt.Println("Value of min after loop :", min)
}
