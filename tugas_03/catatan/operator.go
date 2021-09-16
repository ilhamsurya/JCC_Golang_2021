package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	soal_1()

	// soal 2
	soal_2()

	// soal 3
	soal_3()

	// soal 4
	soal_4()

	// soal 5
	soal_5()
}

func arithmetic() {

	// operator penjumlahan
	jumlah := 8 + 3
	fmt.Println(jumlah) // hasilnya 13

	// operator pengurangan
	kurang := 8 - 3
	fmt.Println(kurang) // hasilnya 5

	// operator perkalian
	kali := 8 * 3
	fmt.Println(kali) // hasilnya 24

	// operator pembagian
	bagi := 8 / 4
	fmt.Println(bagi) // hasilnya 2

	// operator modulus
	modulus := 8 % 3
	fmt.Println(modulus) // hasilnya 2
}

func comparison_operator() {
	var angka = 8

	fmt.Println(angka > 5) // true

	fmt.Println(angka < 5) // false

	fmt.Println(angka >= 5) // true

	fmt.Println(angka <= 5) // false

	fmt.Println(angka == 5) // false

	fmt.Println(angka != 5) // true
}

func conditional() {
	var minimarketStatus = "close"
	var minuteRemainingToOpen = 5
	if minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("minimarket buka sebentar lagi, saya tungguin")
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}
}

func soal_4() {
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var num1, err = strconv.Atoi(angkaPertama)
	var num2, err1 = strconv.Atoi(angkaKedua)
	var num3, err2 = strconv.Atoi(angkaKetiga)
	var num4, err3 = strconv.Atoi(angkaKeempat)

	if err == nil || err1 == nil || err2 == nil || err3 == nil {
		var jumlah = num1 + num2 + num3 + num4
		fmt.Println(jumlah) // 124
	}
}

func soal_5() {
	kalimat := "halo halo bandung"
	angka := 2021

	find := "halo"
	replaceWith := "Hi"

	kalimat = strings.Replace(kalimat, find, replaceWith, 2)
	str_angka := strconv.Itoa(angka)

	fmt.Println(`"` + kalimat + `"` + "-" + str_angka)
}
