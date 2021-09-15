package main

import (
	"fmt"
	"strconv"
)

type hitung interface {
	perkalian() int
	pembagian() float64
	penjumlahan()
}

func main() {
	// soal 1
var numberA int = 4
var numberB *int = &numberA

fmt.Println("numberA (value)   :", numberA)
fmt.Println("numberA (address) :", &numberA)
fmt.Println("numberB (value)   :", *numberB)
fmt.Println("numberB (address) :", numberB)

fmt.Println("")

numberB = 5

fmt.Println("numberA (value)   :", numberA)
fmt.Println("numberA (address) :", &numberA)
fmt.Println("numberB (value)   :", *numberB)
fmt.Println("numberB (address) :", numberB)
	soal_2()
	strconv.Itoa("test")

}
func soal_2() {
	angka := 20
	var angkaBaru *int = &angka
	*angkaBaru += 5
	angka = 20 - *angkaBaru
	angka += 3
	*angkaBaru += 20
	fmt.Println(angka)
}
