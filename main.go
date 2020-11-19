package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)

	fmt.Println(!true)
}

/*
Expression Value
true && true -> true
true && false -> false
false && true -> false
false && false -> false
*/

/*
Expression Value
true || true -> true
true || false -> true
false || true -> true
false || false -> false
*/

/*
Expression Value
!true -> false
!false -> true
*/
