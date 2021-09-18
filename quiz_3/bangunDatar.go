package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/bangun-datar/segitiga-sama-sisi", func(res http.ResponseWriter, req *http.Request) {
		hitungSegitigaSamaSisi(res, req)
	})
	http.HandleFunc("/bangun-datar/persegi", func(res http.ResponseWriter, req *http.Request) {
		hitungPersegi(res, req)
	})
	http.HandleFunc("/bangun-datar/persegi-panjang", func(res http.ResponseWriter, req *http.Request) {
		hitungPersegiPanjang(res, req)
	})
	http.HandleFunc("/bangun-datar/lingkaran", func(res http.ResponseWriter, req *http.Request) {
		hitungLingkaran(res, req)
	})
	http.HandleFunc("/bangun-datar/jajar-genjang", func(res http.ResponseWriter, req *http.Request) {
		hitungJajarGenjang(res, req)
	})

	http.ListenAndServe(":8080", nil)
}

func hitungSegitigaSamaSisi(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("hitung") == "luas" {
		getAlas := req.FormValue("alas")
		getTinggi := req.FormValue("tinggi")

		alas, err := strconv.Atoi(getAlas)
		tinggi, err2 := strconv.Atoi(getTinggi)
		if err == nil || err2 == nil {
			fmt.Printf("%T \n %v", alas, tinggi)
		}
		luasSegitiga := alas * tinggi / 2

		fmt.Fprintln(res, "Ini perhitungan luas segitiga sama sisi")
		fmt.Fprintln(res, luasSegitiga)
	} else if req.FormValue("hitung") == "keliling" {
		getAlas := req.FormValue("alas")
		alas, err := strconv.Atoi(getAlas)
		if err == nil {
			fmt.Printf("%T \n %v", alas)
		}
		kelilingSegitiga := 3 * alas

		fmt.Fprintln(res, "Ini perhitungan keliling segitiga sama sisi")
		fmt.Fprintln(res, kelilingSegitiga)
	}
}

func hitungPersegi(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("hitung") == "luas" {
		getSisi := req.FormValue("sisi")

		sisi, err := strconv.Atoi(getSisi)

		if err == nil {
			fmt.Printf("%T \n %v", sisi)
		}
		luasPersegi := sisi * sisi

		fmt.Fprintln(res, "Ini perhitungan luas persegi")
		fmt.Fprintln(res, luasPersegi)
	} else if req.FormValue("hitung") == "keliling" {
		getSisi := req.FormValue("sisi")

		sisi, err := strconv.Atoi(getSisi)

		if err == nil {
			fmt.Printf("%T \n %v", sisi)
		}
		kelilingPersegi := 4 * sisi

		fmt.Fprintln(res, "Ini perhitungan keliling persegi")
		fmt.Fprintln(res, kelilingPersegi)
	}
}

func hitungPersegiPanjang(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("hitung") == "luas" {
		getPanjang := req.FormValue("panjang")
		getLebar := req.FormValue("lebar")
		panjang, err := strconv.Atoi(getPanjang)
		lebar, err2 := strconv.Atoi(getLebar)
		if err == nil || err2 == nil {
			fmt.Printf("%T \n %v", panjang, lebar)
		}
		luasPersegiPanjang := panjang * lebar

		fmt.Fprintln(res, "Ini perhitungan luas persegi panjang")
		fmt.Fprintln(res, luasPersegiPanjang)
	} else if req.FormValue("hitung") == "keliling" {
		getPanjang := req.FormValue("panjang")
		getLebar := req.FormValue("lebar")
		panjang, err := strconv.Atoi(getPanjang)
		lebar, err2 := strconv.Atoi(getLebar)
		if err == nil || err2 == nil {
			fmt.Printf("%T \n %v", panjang, lebar)
		}
		kelilingPersegiPanjang := (2 * panjang) + (2 * lebar)

		fmt.Fprintln(res, "Ini perhitungan keliling persegi panjang ")
		fmt.Fprintln(res, kelilingPersegiPanjang)
	}
}

func hitungLingkaran(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("hitung") == "luas" {
		getJariJari := req.FormValue("jarijari")

		jariJari, err := strconv.Atoi(getJariJari)

		if err == nil {
			fmt.Printf("%T \n %v", jariJari)
		}
		luasLingkaran := math.Pi * float64(jariJari*jariJari)

		fmt.Fprintln(res, "Ini perhitungan luas lingkaran")
		fmt.Fprintln(res, luasLingkaran)
	} else if req.FormValue("hitung") == "keliling" {
		getJariJari := req.FormValue("jarijari")

		jariJari, err := strconv.Atoi(getJariJari)

		if err == nil {
			fmt.Printf("%T \n %v", jariJari)
		}
		kelilingLingkaran := math.Pi * float64(jariJari+jariJari)

		fmt.Fprintln(res, "Ini perhitungan keliling lingkaran ")
		fmt.Fprintln(res, kelilingLingkaran)
	}
}

func hitungJajarGenjang(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("hitung") == "luas" {
		getAlas := req.FormValue("alas")
		getTinggi := req.FormValue("tinggi")
		alas, err := strconv.Atoi(getAlas)
		tinggi, err2 := strconv.Atoi(getTinggi)
		if err == nil || err2 == nil {
			fmt.Printf("%T \n %v", alas, tinggi)
		}
		luasJajarGenjang := alas * tinggi

		fmt.Fprintln(res, "Ini perhitungan luas jajar genjang")
		fmt.Fprintln(res, luasJajarGenjang)
	} else if req.FormValue("hitung") == "keliling" {
		getAlas := req.FormValue("alas")
		getSisi := req.FormValue("sisi")
		alas, err := strconv.Atoi(getAlas)
		sisi, err2 := strconv.Atoi(getSisi)
		if err == nil || err2 == nil {
			fmt.Printf("%T \n %v", alas, sisi)
		}
		kelilingJajarGenjang := (2 * alas) + (2 * sisi)

		fmt.Fprintln(res, "Ini perhitungan keliling jajar genjang ")
		fmt.Fprintln(res, kelilingJajarGenjang)
	}
}
