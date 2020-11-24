package main

import "fmt"

func main() {
	/*
		we just wrote so that instead of just printing the numbers
		1–10 on each line, it also specifies whether or not the number is
		even or odd: */

	var i int16
	i = 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i, " ", "even")
		} else {
			fmt.Println(i, " ", "odd")
		}

		i++
	}
}
