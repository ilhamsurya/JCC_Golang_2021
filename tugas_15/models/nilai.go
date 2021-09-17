package models

import (
	"time"
)

type (
	// Movie
	NilaiMahasiswa struct {
		ID          int       `json:"id"`
		Nama        string    `json:"nama"`
		MataKuliah  string    `json:"mataKuliah"`
		IndeksNilai string    `json:"indeks"`
		Nilai       int       `json:"nilai"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
