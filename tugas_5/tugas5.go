package main

import (
	"fmt"
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
	fmt.Println("=========== Soal 1 =============")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, tinggi)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("luas:", luas)
	fmt.Println("keliling:", keliling)
	fmt.Println("volume:", volume)
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2*panjang + 2*lebar
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

func soal_2() {
	fmt.Println("=========== Soal 2 =============")
	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")

	fmt.Println(john)
	fmt.Println(sarah)
}

func introduce(name string, gender string, job string, age string) (identity string) {
	if gender == "laki-laki" {
		identity = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else if gender == "perempuan" {
		identity = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}
	return identity
}

func soal_3() {
	fmt.Println("=========== Soal 3 =============")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavorit("John", buah...)
	fmt.Println("")
}

func buahFavorit(name string, buah ...string) {
	fmt.Print("Halo nama saya ", name)
	fmt.Print(" dan buah favorit saya adalah: ")
	for _, satu := range buah {
		fmt.Print(satu, ",")
	}
}

func soal_4() {
	fmt.Println("=========== Soal 4 =============")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	var tambahDataFilm = func(nama string, durasi string, genre string, year string) {
		var data = map[string]string{}
		data["genre"] = genre
		data["jam"] = durasi
		data["tahun"] = year
		data["title"] = nama

		dataFilm = append(dataFilm, data)

	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
