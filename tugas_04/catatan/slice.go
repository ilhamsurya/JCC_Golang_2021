package main

import (
	"fmt"
)

func main() {
	//! difference between slice and array is that when initialized slice it will be undefined []
	//! but array will have some sort of element like [...] or [3]
	//normal slice
	var slice = []string{"test", "test1", "test2"}
	fmt.Println(slice)

	// * second way to use slice, is to accessing some element in array ex:
	var arr = [...]int{1, 2, 3, 4, 5}
	var slice_arr = arr[0:2]
	fmt.Println(slice_arr)

	// ? func cap() used to count slice capacity
	fmt.Println(cap(slice_arr))

	// ? func append() used to adding element in slice
	slice_arr = append(slice_arr, 9)
	fmt.Println(slice_arr)

	// ? func copy() used to copy element from slice
	// ? there are 2 params dst and src

	dst := make([]int, 3)
	src := []int{1, 3, 5, 7, 9}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
