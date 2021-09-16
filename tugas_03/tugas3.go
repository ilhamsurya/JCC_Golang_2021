package main

import (
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

	// soal 4
	soal_4()
}

func soal_1() {

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjang, err = strconv.Atoi(panjangPersegiPanjang)
	var lebar, err1 = strconv.Atoi(lebarPersegiPanjang)
	var alas, err2 = strconv.Atoi(alasSegitiga)
	var tinggi, err3 = strconv.Atoi(tinggiSegitiga)

	var kelilingPersegiPanjang = 2*panjang + 2*lebar
	var luasSegitiga = alas * tinggi / 2

	if err == nil || err1 == nil || err2 == nil || err3 == nil {
		var keteranganPersegiPanjang = fmt.Sprintf("%s%d", "Keliling persegi panjang adalah = ", kelilingPersegiPanjang)
		var keteranganLuasSegitiga = fmt.Sprintf("%s%d", "luas segitiga adalah = ", luasSegitiga)
		fmt.Println(keteranganPersegiPanjang)
		fmt.Println(keteranganLuasSegitiga)
	}

}

func soal_2() {

	var nilaiJohn = 80
	var nilaiDoe = 50

	//check nilai john
	if nilaiJohn >= 80 {
		fmt.Println("Index John: A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Index John: B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Index John: C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Index John: D")
	} else if nilaiJohn < 50 {
		fmt.Println("Index John: E")
	} else {
		fmt.Println("Undefined")
	}
	//check nilai doe
	if nilaiDoe >= 80 {
		fmt.Println("Index Doe: A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Index Doe: B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Index Doe: C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Index Doe: D")
	} else if nilaiDoe < 50 {
		fmt.Println("Index Doe: E")
	} else {
		fmt.Println("Undefined")
	}
}

func soal_3() {
	var tanggal = 27
	var bulan = 10
	var tahun = 2000

	switch {
	case tanggal == 27 && bulan == 10 && tahun == 2000:
		fmt.Println("tanggal lahir anda 27/10/2000")
	default:
		fmt.Println("Harap masukan tanggal lahir")
	}

}

func soal_4() {
	var tahunLahir = 2000
	switch {
	case tahunLahir >= 1944 && tahunLahir <= 1964:
		fmt.Println("Anda termasuk ke golongan Baby Boomer")
	case tahunLahir >= 1965 && tahunLahir <= 1979:
		fmt.Println("Anda termasuk ke golongan Generasi X")
	case tahunLahir >= 1980 && tahunLahir <= 1994:
		fmt.Println("Anda termasuk ke golongan Generasi Y")
	case tahunLahir >= 1995 && tahunLahir <= 2015:
		fmt.Println("Anda termasuk ke golongan Generasi Z")
	default:
		fmt.Println("Golongan anda tidak diketahui")
	}
}
