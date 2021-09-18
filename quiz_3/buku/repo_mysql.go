package buku

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"quiz3/config"
	"quiz3/models"
	"time"
)

const (
	table          = "buku"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Buku, error) {
	var daftarBuku []models.Buku
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
		var buku models.Buku
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&buku.ID,
			&buku.Title,
			&buku.Description,
			&buku.ImageUrl,
			&buku.ReleaseYear,
			&buku.Price,
			&buku.TotalPage,
			&buku.KategoriKetebalan,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		buku.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		daftarBuku = append(daftarBuku, buku)
	}
	return daftarBuku, nil
}

func FilterBook(ctx context.Context, title string, minYear int, maxYear int) ([]models.Buku, error) {
	var daftarBuku []models.Buku
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v WHERE title = '%v', release_year > %v AND release_year < %v Order By created_at DESC", table,
		title,
		minYear,
		maxYear,
	)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var buku models.Buku
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&buku.ID,
			&buku.Title,
			&buku.Description,
			&buku.ImageUrl,
			&buku.ReleaseYear,
			&buku.Price,
			&buku.TotalPage,
			&buku.KategoriKetebalan,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		if err != nil {
			log.Fatal(err)
		}

		buku.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		daftarBuku = append(daftarBuku, buku)
	}
	return daftarBuku, nil
}

// Insert Movie
func Insert(ctx context.Context, buku models.Buku) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, kategori_ketebalan, updated_at) values ('%v','%v','%v',%v,'%v','%v', '%v',NOW())",
		table,
		buku.Title,
		buku.Description,
		buku.ImageUrl,
		buku.ReleaseYear,
		buku.Price,
		buku.TotalPage,
		buku.KategoriKetebalan,
	)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}
func Update(ctx context.Context, buku models.Buku, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title ='%s', description = '%s', image_url = '%s', release_year = %d, price = '%s',total_page = '%s',updated_at = NOW() where id = %s",
		table,
		buku.Title,
		buku.Description,
		buku.ImageUrl,
		buku.ReleaseYear,
		buku.Price,
		buku.TotalPage,
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
