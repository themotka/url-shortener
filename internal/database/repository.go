package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetKeyByValue(value string) (string, error) {
	var key string
	query := "SELECT key FROM hashTable WHERE value = $1"
	err := r.db.QueryRow(query, value).Scan(&key)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("value not found")
		}
		return "", err
	}
	return key, nil
}

func (r *Repository) GetValueByKey(key string) (string, error) {
	var value string
	query := "SELECT value FROM hashTable WHERE key = $1"
	err := r.db.QueryRow(query, key).Scan(&value)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "", fmt.Errorf("key not found")
		}
		return "", err
	}
	return value, nil
}

func (r *Repository) Put(key string, value string) error {
	query := "INSERT INTO hashTable (key, value) VALUES ($1, $2)"
	_, err := r.db.Exec(query, key, value)
	if err != nil {
		return err
	}
	return nil
}
