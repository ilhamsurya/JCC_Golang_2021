package main

import (
	"fmt"
)

func main() {
	//normal array (total element defined)
	var arr = [3]string{"test", "test1", "test2"}
	fmt.Println(arr)

	//normal array (total element undefined)
	var arr_undefined = [...]int{1,2,3,4,5}
	fmt.Println(arr_undefined)
	fmt.Println("check element \t:", len(arr_undefined))
	
	//array multidimensional
	//? first value defined the array size
	//? second value defined the multidimensional size
	var arr_multidimensional = [2][2]int{[2]string{"test", "test2"}, [2]{"hai", "hai2"}
	fmt.Println(arr_multidimensional)

}

}