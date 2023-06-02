package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

const addUrl = `
  INSERT INTO url(original, short) VALUES(?, ?)
  `

// Create stores the original url and short url string in the database.
func (r *repository) Create(ctx context.Context, original string, short string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		log.Fatalln("couldnt commit, ", err)
	}

	defer func() {
		if err := tx.Commit(); err != nil {
			log.Fatalln(err)
		}
	}()

	_, err = tx.Exec(addUrl, original, short)
	if err != nil {
		fmt.Println("we got an error yo")
		return err
	}

	return nil
}

const getUrl = `
  SELECT original FROM url WHERE short=?
  `

// Get queries the database for a long url based on the short url provided.
func (r *repository) Get(ctx context.Context, url string) (string, error) {
	var original string

	row := r.DB.QueryRow(getUrl, url)
	if err := row.Scan(&original); err != nil {
		return "", err
	}
	return original, nil
}

const getShortFromOriginal = `
  SELECT short from URL where original=?
  `

// GetExistingShortUrl checks the database to see if a short url already exists for a given long url
// this is to prevent duplication of long url's in the database
func (r *repository) GetExistingShortUrl(ctx context.Context, original string) (string, error) {
	var short string

	row := r.DB.QueryRow(getShortFromOriginal, original)
	err := row.Scan(&short)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", err
		}
	}
	return short, nil
}
