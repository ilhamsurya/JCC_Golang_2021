package main

import (
	"fmt"
)

func main() {
	//? Map is some kind of slice that has their own unique key
	//normal map
	var food map[string]int
	food = map[string]int{}

	food["nasi goreng"] = 15000
	food["sate ayam"] = 30000

	fmt.Println("nasi goreng", food["nasi goreng"])

	// * intialized map (best practice)
	var drink = map[string]int{
		"jus": 20000,
		"teh": 5000,
	}
	fmt.Println(drink["jus"])

	// * deleting map element using delete()
	delete(drink, "teh")
	fmt.Println(drink)

	// * detecting element from its key
	var value, isExist = drink["jus"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}
}
