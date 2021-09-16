package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// soal 1ss
	soal_1()
	// soal 2
	soal_2()
	// soal 3
	soal_3()

}

func printJCC(kalimat string, tahun int) {
	fmt.Println(kalimat, tahun)
}

func soal_1() {
	fmt.Println("=========== Soal 1 =============")
	jcc := "Candradimuka Jabar Coding Camp"
	year := 2021
	defer printJCC(jcc, year)
	fmt.Println("Run Application")
}

func soal_2() {
	fmt.Println("=========== Soal 2 =============")
	defer kelilingSegitigaSamaSisi(4, true)
	defer kelilingSegitigaSamaSisi(8, false)
	defer kelilingSegitigaSamaSisi(0, true)
	err := kelilingSegitigaSamaSisi(0, false)

	if err != nil {
		defer recoveryFunction()
		panic("Anda belum menginput sisi dari segitiga sama sisi")
	}

}
func kelilingSegitigaSamaSisi(sisi int, stats bool) error {
	var infoKeliling string
	s := strconv.Itoa(sisi)
	kelilingSegitga := 3 * sisi
	l := strconv.Itoa(kelilingSegitga)
	if sisi > 0 && stats == true {
		infoKeliling = "keliling segitiga sama sisinya dengan sisi " + s + " adalah " + l
		fmt.Println(infoKeliling)
	} else if sisi > 0 && stats == false {
		infoKeliling = l
		fmt.Println(infoKeliling)
	} else if sisi == 0 && stats == true {
		return errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if sisi == 0 && stats == false {
		return errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else {
		return errors.New("terjadi error")
	}
	return nil
}

func recoveryFunction() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func tambahAngka(angka1 int, angka2 *int) {
	*angka2 = angka1 + *angka2
	fmt.Println(*angka2)
}

func soal_3() {
	fmt.Println("============ Soal 3 ==============")
	angka := 1
	defer tambahAngka(7, &angka)
	defer tambahAngka(6, &angka)
	defer tambahAngka(-1, &angka)
	defer tambahAngka(9, &angka)
	fmt.Println("Berhasil menjumlahkan angka:")
}
