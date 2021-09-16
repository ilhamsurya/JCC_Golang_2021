package main

import (
	"fmt"
)

func main() {
	// * looping using for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// * for using range in array
	var nums = [4]int{1, 3, 5, 6}

	for i, num := range nums {
		fmt.Printf("elemen %d : %s\n", i, num)
	}

}
