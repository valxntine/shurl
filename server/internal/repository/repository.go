package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	DB *sql.DB
}

const createTable = `
  CREATE TABLE IF NOT EXISTS url (
	id INTEGER NOT NULL PRIMARY KEY,
	original TEXT,
	short TEXT
)
  `

// NewRepository returns a repository struct to allow calls to the database to be made
func NewRepository() (*repository, error) {
	conn, err := sql.Open("sqlite3", "urls.db")
	if err != nil {
		fmt.Println(fmt.Errorf("shit the bed in the constructor %w", err))
		panic(err)
	}

	_, err = conn.Exec(createTable)
	if err != nil {
		fmt.Println("shit the bed creating table")
		return nil, err
	}

	return &repository{
		DB: conn,
	}, nil
}
