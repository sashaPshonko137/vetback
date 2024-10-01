package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"vetback/internal/repo"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(dbPath string) (repo.Repo, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Storage{db: db}, nil
}
