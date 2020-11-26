/*
Slice
*/

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)

	source := []string{"A", "B", "C"}
	destination := make([]string, 2)

	copy(destination, source)
	fmt.Println(source, destination)
}

/*
After running this program slice1 has [1,2,3] and slice2 has [1,2].
The contents of slice1 are copied into slice2, but because slice2 has
room for only two elements, only the first two elements of slice1 are
copied.
*/
