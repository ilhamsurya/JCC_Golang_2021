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

func soal_1() {

	var jabar = "Jabar "
	var coding = "Coding "
	var camp = "Camp "
	var golang = "Golang "

	fmt.Println(jabar + coding + camp + golang)
}

func soal_2() {
	halo := "Halo Dunia"
	var find = "Dunia"
	var replaceWith = "Golang"

	var newHalo = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newHalo)
}

func soal_3() {
	var kataPertama = "saya "
	var kataKedua = "senang "
	var kataKetiga = "belajar "
	var kataKeempat = "golang "

	kataKedua = strings.Title(kataKedua)
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Println(kataPertama + kataKedua + kataKetiga + kataKeempat)
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
