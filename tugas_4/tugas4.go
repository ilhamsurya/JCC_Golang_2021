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

	// soal 5
	soal_5()
}

func soal_1() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%2 != 0 {
			fmt.Println("I Love Coding")
		} else if i%2 == 0 {
			fmt.Println("Candradimuka")
		} else if i%2 != 0 {
			fmt.Println("JCC")
		}
	}
}

func soal_2() {
	for i := 0; i < 6; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}
}

func soal_3() {
	var kalimat = []string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	i := 0

	for j := 0; j < 2; j++ {
		copy(kalimat[i:], kalimat[i+1:])
		kalimat[len(kalimat)-1] = ""
		kalimat = kalimat[:len(kalimat)-1]
	}

	fmt.Println(kalimat)
}

func soal_4() {
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i := 1; i < len(sayuran); i++ {
		fmt.Println(i, sayuran[i])
	}
}

func soal_5() {
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for i, element := range satuan {
		fmt.Println(i, "=", element)
	}

	satuan["Volume Balok"] = 168
	fmt.Println("Volume Balok", "=", satuan["Volume Balok"])

}
