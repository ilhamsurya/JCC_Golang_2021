// struct
package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}
type person struct {
	name string
	age  int
}

func main() {
	// normal struct
	structStudent()

	// struct Literals
	structLiterals()

	// Embedded struct

}

func structStudent() {
	var john student
	john.name = "john doe"
	john.grade = 2

	fmt.Println("name  :", john.name)
	fmt.Println("grade :", john.grade)
}

func structLiterals() {
	// cara pertama
	var john = student{}
	john.name = "wick"
	john.grade = 2
	// cara kedua tetapi isinya harus berurutan
	var doe = student{"doe", 2}
	// cara ketiga dengan nama property tetapi tidak harus berurutan
	var jack = student{name: "jack", grade: 2}

	fmt.Println("student 1 :", john.name)
	fmt.Println("student 2 :", doe.name)
	fmt.Println("student 3 :", jack.name)
}

func embeddedStruct() {

}
