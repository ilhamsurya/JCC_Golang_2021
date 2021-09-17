package nilai

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"tugas15/config"
	"tugas15/models"
)

const (
	table          = "nilai_mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {
	var daftarNilai []models.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.NilaiMahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Nama,
			&nilai.MataKuliah,
			&nilai.IndeksNilai,
			&nilai.Nilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		daftarNilai = append(daftarNilai, nilai)
	}
	return daftarNilai, nil
}

// Insert Movie
func Insert(ctx context.Context, nilaiMhs models.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, indeks_nilai, nilai, updated_at) values('%v','%v','%v',%v, NOW())", table,
		nilaiMhs.Nama,
		nilaiMhs.MataKuliah,
		nilaiMhs.IndeksNilai,
		nilaiMhs.Nilai)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
func Update(ctx context.Context, nilaiMhs models.NilaiMahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama ='%s', mata_kuliah = '%s', indeks_nilai = '%s', nilai = %d, updated_at = NOW() where id = %s",
		table,
		nilaiMhs.Nama,
		nilaiMhs.MataKuliah,
		nilaiMhs.IndeksNilai,
		nilaiMhs.Nilai,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
