package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

type lingkaran struct {
	jariJari float64
}

type hitungBangunDatar interface {
	luas() float64
	keliling() float64
}
type hitungBangunRuang interface {
	volume() float64
}
type tabung struct {
	jariJari, tinggi float64
}

type persegiPanjang struct {
	panjang, lebar int
}

func main() {
	// soal 1ss
	soal_1()
	// soal 3
	soal_3()
	// soal 3
	soal_4()
	// soal 2
	soal_2()

}

func soal_1() {
	fmt.Println("=========== Soal 1 =============")
	var wg sync.WaitGroup
	var phones = []string{}
	var phoneName = []string{"Xiaomi", "Asus", "A", "Iphone", "Samsung",
		"Oppo", "Realme", "Vivo"}
	addPhones(&phones, phoneName...)
	wg.Add(1)
	phoneSort := sortStrings(phones, &wg)
	wg.Add(1)
	go printPhones(phoneSort, &wg)
	wg.Wait()

}

func sortStrings(phones []string, wg *sync.WaitGroup) []string {
	sort.Strings(phones)
	wg.Done()
	return phones
}

func printPhones(phones []string, wg *sync.WaitGroup) {
	for i := 1; i < len(phones); i++ {
		time.Sleep(time.Second * 1)
		fmt.Println(i, phones[i])
	}
	wg.Done()
}

func addPhones(phones *[]string, name ...string) {
	for _, phone := range name {
		*phones = append(*phones, phone)
	}
}

func soal_2() {
	fmt.Println("=========== Soal 2 =============")
	var wg sync.WaitGroup
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, &wg, movies...)

	for value := range moviesChannel {
		fmt.Println(value)

	}
	wg.Done()
}

func getMovies(movie chan string, wg *sync.WaitGroup, movies ...string) {
	wg.Add(1)
	for _, y := range movies {
		go func(y string) {
			movie <- y
		}(y)

	}
	wg.Wait()

}

func soal_3() {
	fmt.Println("============ Soal 3 ==============")
	var lingkaran1, lingkaran2, lingkaran3 hitungBangunDatar
	var tabung1, tabung2, tabung3 hitungBangunRuang
	var wg sync.WaitGroup
	lingkaran1 = lingkaran{8}
	lingkaran2 = lingkaran{14}
	lingkaran3 = lingkaran{20}
	tabung1 = tabung{8, 10}
	tabung2 = tabung{14, 10}
	tabung3 = tabung{20, 10}
	fmt.Println("== Lingkaran 1==")
	hitungLingkaran(lingkaran1, &wg)
	fmt.Println("== Lingkaran 2 ==")
	hitungLingkaran(lingkaran2, &wg)
	fmt.Println("== Lingkaran 3 ==")
	hitungLingkaran(lingkaran3, &wg)
	fmt.Println("== Tabung 1 ==")
	hitungTabung(tabung1, &wg)
	fmt.Println("== Tabung 2 ==")
	hitungTabung(tabung2, &wg)
	fmt.Println("== Tabung 3 ==")
	hitungTabung(tabung3, &wg)
	wg.Wait()
}

func hitungLingkaran(bangunLingkaran hitungBangunDatar, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Luas lingkaran adalah: ", bangunLingkaran.luas())
	fmt.Println("Keliling lingkaran adalah: ", bangunLingkaran.keliling())
	wg.Done()
}

func hitungTabung(bangunTabung hitungBangunRuang, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Volume tabung adalah: ", bangunTabung.volume())
	wg.Done()
}
func (l lingkaran) luas() float64 {
	return math.Round(math.Pi * math.Pow(l.jariJari, 2))

}
func (l lingkaran) keliling() float64 {
	return math.Round(math.Pi * (l.jariJari + l.jariJari))
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * 2 * t.tinggi
}

func soal_4() {
	fmt.Println("============ Soal 4 ==============")
	var ch1 = make(chan int)
	go luasPersegiPanjang(10, 5, ch1)

	var ch2 = make(chan int)
	go kelilingPersegiPanjang(12, 6, ch2)

	var ch3 = make(chan float64)
	go volumeTabung(3.14, 12.5, ch3)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-ch1:
			fmt.Printf("Luas \t: %d \n", luas)
		case keliling := <-ch2:
			fmt.Printf("Keliling \t: %d \n", keliling)
		case volume := <-ch3:
			fmt.Printf("volume \t: %.2f \n", volume)
		}
	}

}

func luasPersegiPanjang(panjang int, lebar int, ch chan int) {
	ch <- panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int, ch chan int) {
	ch <- (2 * panjang) + (2 * lebar)
}

func volumeTabung(jariJari float64, tinggi float64, ch chan float64) {
	ch <- math.Pi * jariJari * 2 * tinggi
}
