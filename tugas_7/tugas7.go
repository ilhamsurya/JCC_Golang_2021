package main

import (
	"fmt"
)

type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}
type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title, genre   string
	year, duration int
}

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
	var nanas = buah{"Nanas", "Kuning", false, 9000}
	var jeruk = buah{"Jeruk", "Oranye", true, 8000}
	var semangka = buah{"Semangka", "Hijau & Merah", true, 10000}
	var pisang = buah{"Pisang", "Kuning", false, 5000}

	fmt.Println("Nama ", "Warna ", "Biji ", "Harga ")
	if nanas.adaBijinya == true {
		biji := "Ada"
		fmt.Println(nanas.nama, " ", nanas.warna, " ", biji, " ", nanas.harga)
	} else {
		biji := "Tidak"
		fmt.Println(nanas.nama, " ", nanas.warna, " ", biji, " ", nanas.harga)
	}
	if jeruk.adaBijinya == true {
		biji := "Ada"
		fmt.Println(jeruk.nama, " ", jeruk.warna, " ", biji, " ", jeruk.harga)
	} else {
		biji := "Tidak"
		fmt.Println(jeruk.nama, " ", jeruk.warna, " ", biji, " ", jeruk.harga)
	}
	if pisang.adaBijinya == true {
		biji := "Ada"
		fmt.Println(pisang.nama, " ", pisang.warna, " ", biji, " ", pisang.harga)
	} else {
		biji := "Tidak"
		fmt.Println(pisang.nama, " ", pisang.warna, " ", biji, " ", pisang.harga)
	}
	if semangka.adaBijinya == true {
		biji := "Ada"
		fmt.Println(semangka.nama, " ", semangka.warna, " ", biji, " ", semangka.harga)
	} else {
		biji := "Tidak"
		fmt.Println(semangka.nama, " ", semangka.warna, " ", biji, " ", semangka.harga)
	}
}

func soal_2() {
	fmt.Println("=========== Soal 2 =============")
	var bangunPersegi = persegi{6}
	var bangunPersegiPanjang = persegiPanjang{6, 10}
	var bangunSegitiga = segitiga{10, 4}

	bangunPersegi.luasPersegi()
	bangunPersegiPanjang.luasPersegiPanjang()
	bangunSegitiga.luasSegitiga()
}

func (p persegi) luasPersegi() {
	luas := p.sisi * p.sisi
	fmt.Println("Luas persegi adalah: ", luas)
}
func (s segitiga) luasSegitiga() {
	luas := s.alas * s.tinggi / 2
	fmt.Println("Luas segitiga adalah: ", luas)
}
func (pp persegiPanjang) luasPersegiPanjang() {
	luas := pp.panjang * pp.lebar
	fmt.Println("Luas persegi panjang adalah: ", luas)
}
func soal_3() {
	fmt.Println("=========== Soal 3 =============")
	var color []string
	var color_list = []string{"merah", "hijau", "kuning"}
	xiaomi := phone{"Redmi10", "Xiaomi", 2021, color}
	color = phone.addColors(xiaomi, color_list...)
	xiaomi.colors = color
	fmt.Println(xiaomi)
}

func (p phone) addColors(warna ...string) []string {
	var color_palette []string
	for _, item := range warna {
		color_palette = append(color_palette, item)
	}
	return color_palette
}

func soal_4() {
	fmt.Println("=========== Soal 4 =============")
	var dataFilm = []map[string]string{}
	var lotr, avenger, spiderman, juon movie

	movie.tambahDataFilm(lotr, &dataFilm, "LOTR", "2 jam", "action", "1999")
	movie.tambahDataFilm(avenger, &dataFilm, "avenger", "2 jam", "action", "2019")
	movie.tambahDataFilm(spiderman, &dataFilm, "spiderman", "2 jam", "action", "2004")
	movie.tambahDataFilm(juon, &dataFilm, "juon", "2 jam", "horror", "2004")

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

func (m movie) tambahDataFilm(dataFilm *[]map[string]string, nama string, durasi string, genre string, year string) {
	var data = map[string]string{}
	data["genre"] = genre
	data["duration"] = durasi
	data["year"] = year
	data["title"] = nama
	*dataFilm = append(*dataFilm, data)
}
