package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"quiz3/buku"
	"quiz3/config"
	"quiz3/models"
	"quiz3/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/paimanbandi/rupiah"
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
	router.GET("/books_list", GetBuku)
	router.POST("/books/create", PostBuku)
	router.PUT("/books/:id/update", UpdateBuku)
	router.DELETE("/books/:id/delete", DeleteBuku)
	http.HandleFunc("/books", func(res http.ResponseWriter, req *http.Request) {
		FilterBuku(res, req)
	})

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetBuku(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nilai, err := buku.GetAll(ctx)

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
func PostBuku(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var dftBuku models.Buku
	bukuPrice := dftBuku.Price
	price, err := strconv.Atoi(bukuPrice)
	if err == nil {
		fmt.Printf("%T \n %v", price)
	}
	dftBuku.Price = rupiah.FormatFloat64ToRp(float64(price))
	// if dftBuku.ReleaseYear < 1980 || dftBuku.ReleaseYear > 2021 {
	// 	defer recoveryFunction()
	// 	panic("Tahun tidak boleh kurang dari 1980 dan lebih dari 2021")
	// }
	if isValidUrl(dftBuku.ImageUrl) {
		defer recoveryFunction()
		panic("URL Tidak valid")
	}
	totalPage := dftBuku.TotalPage
	page, err := strconv.Atoi(totalPage)
	if err == nil {
		fmt.Printf("%T \n %v", page)
	}
	if page <= 100 {
		dftBuku.KategoriKetebalan = "tipis"
	} else if page >= 101 && page <= 200 {
		dftBuku.KategoriKetebalan = "sedang"
	} else if page >= 201 {
		dftBuku.KategoriKetebalan = "tebal"
	} else {
		dftBuku.KategoriKetebalan = "tidak terdefinisi"
	}
	if err := json.NewDecoder(r.Body).Decode(&dftBuku); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}
	if err := buku.Insert(ctx, dftBuku); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
func UpdateBuku(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var dftBuku models.Buku

	if err := json.NewDecoder(r.Body).Decode(&dftBuku); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	if err := buku.Update(ctx, dftBuku, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

func DeleteBuku(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idBuku = ps.ByName("id")
	if err := buku.Delete(ctx, idBuku); err != nil {
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

func FilterBuku(res http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	getMinYear := req.FormValue("minYear")
	getMaxYear := req.FormValue("maxYear")

	year, err := strconv.Atoi(getMinYear)
	maxYear, err := strconv.Atoi(getMaxYear)
	if err == nil {
		fmt.Printf("%T \n %v", year, maxYear)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nilai, err := buku.FilterBook(ctx, title, year, maxYear)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(res, nilai, http.StatusOK)

}
