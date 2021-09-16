package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     int
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

// GetMovies
func getNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nilaiMahasiswa := nilaiNilaiMahasiswa
		dataNilai, err := json.Marshal(nilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nilMahasiswa NilaiMahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&nilMahasiswa); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			nama := r.PostFormValue("nama")
			matkul := r.PostFormValue("matakuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			nilMahasiswa = NilaiMahasiswa{
				Nama:       nama,
				MataKuliah: matkul,
				Nilai:      nilai,
			}
		}
		if nilMahasiswa.Nilai > 100 {
			defer recoveryFunction()
			panic("Nilai tidak bisa lebih besar dari 100")
		}
		if nilMahasiswa.Nilai >= 80 {
			nilMahasiswa.IndeksNilai = "A"
		} else if nilMahasiswa.Nilai >= 70 && nilMahasiswa.Nilai < 80 {
			nilMahasiswa.IndeksNilai = "B"
		} else if nilMahasiswa.Nilai >= 60 && nilMahasiswa.Nilai < 70 {
			nilMahasiswa.IndeksNilai = "C"
		} else if nilMahasiswa.Nilai >= 50 && nilMahasiswa.Nilai < 60 {
			nilMahasiswa.IndeksNilai = "D"
		} else {
			nilMahasiswa.IndeksNilai = "E"
		}
		nilMahasiswa.ID = rand.Intn(100)
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilMahasiswa)

		dataNilai, err := json.Marshal(nilMahasiswa)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.Write(dataNilai)
		return

	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}
func recoveryFunction() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func main() {
	http.HandleFunc("/daftar_nilai", getNilai)
	http.HandleFunc("/input_nilai", PostNilai)
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
