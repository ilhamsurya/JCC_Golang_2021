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
	fmt.Println("=========== Soal 3 =============")
	var luasLigkaran float64
	var kelilingLingkaran float64
	var phi float64 = 3.14
	var radius float64 = 7

	luasLigkaran = hitungLuasLingkaran(phi, radius)
	kelilingLingkaran = hitungKelilingLingkaran(phi, radius)

	fmt.Println("Luas lingkaran (sebelum diubah)", luasLigkaran)
	fmt.Println("Keliling lingkaran (sebelum diubah)", kelilingLingkaran)

	updateValueLingkaran(&radius, 14)
	luasLigkaran = hitungLuasLingkaran(phi, radius)
	kelilingLingkaran = hitungKelilingLingkaran(phi, radius)

	fmt.Println("Luas lingkaran (setelah diubah)", luasLigkaran)
	fmt.Println("Luas lingkaran (setelah diubah)", kelilingLingkaran)
}

func hitungLuasLingkaran(phi float64, radius float64) float64 {
	return phi * radius * radius
}

func hitungKelilingLingkaran(phi float64, radius float64) float64 {
	return phi * (radius + radius)
}
func updateValueLingkaran(original *float64, value float64) {
	*original = value
}

func soal_2() {
	fmt.Println("=========== Soal 2 =============")
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
}

func introduce(sentence *string, name string, gender string, job string, age string) {
	var newSentence string

	if gender == "laki-laki" {
		newSentence = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun "
	} else if gender == "perempuan" {
		newSentence = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else {
		fmt.Println("Harap masukan gender")
	}
	*sentence = newSentence
}

func soal_3() {
	fmt.Println("=========== Soal 3 =============")
	var buah = []string{}
	var daftar_buah = []string{"test", "Jeruk", "Semangka", "Mangga",
		"Strawberry", "Durian", "Manggis", "Alpukat"}
	appendToItself(&buah, daftar_buah...)

	for i := 1; i < len(buah); i++ {
		fmt.Println(i, buah[i])
	}

}

func appendToItself(slice *[]string, slice_list ...string) {
	for _, satu := range slice_list {
		*slice = append(*slice, satu)
	}

}

func soal_4() {
	fmt.Println("=========== Soal 4 =============")
	var dataFilm = []map[string]string{}

	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")

	i := 1
	for _, item := range dataFilm {
		fmt.Println("------------------")
		fmt.Print(i, " ")
		fmt.Println("title:", item["title"])
		fmt.Println(" ", "duration:", item["duration"])
		fmt.Println(" ", "genre:", item["genre"])
		fmt.Println(" ", "year:", item["year"])
		i++
	}
}

func tambahDataFilm(dataFilm *[]map[string]string, nama string, durasi string, genre string, year string) {
	var data = map[string]string{}
	data["genre"] = genre
	data["duration"] = durasi
	data["year"] = year
	data["title"] = nama
	*dataFilm = append(*dataFilm, data)
}
