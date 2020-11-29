/*
Slice
*/

package main

import "fmt"

func main() {

	// make new
	// new - only allocates - no init memory
	// make - allocates and init memory - non zerored storage
	score := make(map[string]int)
	score["faisal"] = 90
	fmt.Println(score)

	score["hitesh"] = 78
	score["rahul"] = 60
	score["nikky"] = 58
	score["ron"] = 98

	for key, value := range score {
		fmt.Printf("Score of %v is %v \n", key, value)
	}

	delete(score, "nikky")
	for key, value := range score {
		fmt.Printf("Score of %v is %v \n", key, value)
	}

	getFaisalScore := score["faisal"]
	fmt.Println(getFaisalScore)
}
