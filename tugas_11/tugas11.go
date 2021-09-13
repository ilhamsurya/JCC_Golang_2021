package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

type lingkaran struct {
	jariJari float64
}

type hitungBangunDatar interface {
	luas() float64
	keliling() float64
}

func main() {
	// soal 1ss
	soal_1()
	// soal 2
	soal_2()
	// soal 3
	soal_3()

}

func soal_1() {
	fmt.Println("=========== Soal 1 =============")
	var phones = []string{}
	var phoneName = []string{"Xiaomi", "Asus", "Iphone", "Samsung",
		"Oppo", "Realme", "Vivo"}
	addPhones(&phones, phoneName...)

	for i := 1; i < len(phones); i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i, phones[i])
	}
}

func addPhones(phones *[]string, name ...string) {
	for _, phone := range name {
		*phones = append(*phones, phone)
	}
}

func soal_2() {
	fmt.Println("=========== Soal 2 =============")
	var lingkaran1, lingkaran2, lingkaran3 hitungBangunDatar
	lingkaran1 = lingkaran{7}
	lingkaran2 = lingkaran{10}
	lingkaran3 = lingkaran{15}

	fmt.Println("Luas lingkaran 1 adalah: ", lingkaran1.luas())
	fmt.Println("Keliling lingkaran 1 adalah: ", lingkaran1.keliling())
	fmt.Println("Luas lingkaran 2 adalah: ", lingkaran2.luas())
	fmt.Println("Keliling lingkaran 2 adalah: ", lingkaran2.keliling())
	fmt.Println("Luas lingkaran 3 adalah: ", lingkaran3.luas())
	fmt.Println("Keliling lingkaran 3 adalah: ", lingkaran3.keliling())
}

func (l lingkaran) luas() float64 {
	return math.Round(math.Pi * math.Pow(l.jariJari, 2))
}
func (l lingkaran) keliling() float64 {
	return math.Round(math.Pi * (l.jariJari + l.jariJari))
}
func soal_3() {
	fmt.Println("============ Soal 3 ==============")
	var panjang = flag.Int64("panjang", 15, " persegiPanjang")
	var lebar = flag.Int64("lebar", 10, " persegiPanjang")
	flag.Parse()
	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang(*panjang, *lebar))
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang(*panjang, *lebar))
}

func luasPersegiPanjang(panjang int64, lebar int64) int64 {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int64, lebar int64) int64 {
	return (2 * panjang) + (2 * lebar)
}
