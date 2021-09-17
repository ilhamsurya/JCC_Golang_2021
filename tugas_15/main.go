package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tugas15/config"
	"tugas15/models"
	"tugas15/nilai"
	"tugas15/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()
	router.GET("/daftar_nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nilai, err := nilai.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilai, http.StatusOK)
}
func recoveryFunction() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilMhs models.NilaiMahasiswa
	if nilMhs.Nilai > 100 {
		defer recoveryFunction()
		panic("Nilai tidak bisa lebih besar dari 100")
	}
	if nilMhs.Nilai >= 80 {
		nilMhs.IndeksNilai = "A"
	} else if nilMhs.Nilai >= 70 && nilMhs.Nilai < 80 {
		nilMhs.IndeksNilai = "B"
	} else if nilMhs.Nilai >= 60 && nilMhs.Nilai < 70 {
		nilMhs.IndeksNilai = "C"
	} else if nilMhs.Nilai >= 50 && nilMhs.Nilai < 60 {
		nilMhs.IndeksNilai = "D"
	} else {
		nilMhs.IndeksNilai = "E"
	}

	if err := json.NewDecoder(r.Body).Decode(&nilMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := nilai.Insert(ctx, nilMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var nilMhs models.NilaiMahasiswa

	if err := json.NewDecoder(r.Body).Decode(&nilMhs); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	if err := nilai.Update(ctx, nilMhs, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idNilai = ps.ByName("id")
	if err := nilai.Delete(ctx, idNilai); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}
