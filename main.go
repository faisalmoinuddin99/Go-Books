package main

import "fmt"

func main() {
	var array_name [5]int
	array_name[4] = 100

	fmt.Print(array_name)
}

/*
You should see this:
[0 0 0 0 100]

x[4] = 100 should be read “set the fifth element of the array x to 100.”
It might seem strange that x[4] represents the fifth element instead
of the fourth, but like strings, arrays are indexed starting from 0.
Arrays are accessed in a similar way. We could change fmt.Println(x)
to fmt.Println(x[4]) and we would get 100.
*/
