package pg

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DataBase struct {
	Db *sql.DB
}

func (d *DataBase) Write(key string, url string) error {
	query := "INSERT INTO hashTable (key, value) VALUES ($1, $2)"
	_, err := d.Db.Exec(query, key, url)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataBase) IsPresentedByKey(key string) (bool, string) {
	var value string
	query := "SELECT value FROM hashTable WHERE key = $1"
	err := d.Db.QueryRow(query, key).Scan(&value)
	if err != nil {
		return false, value
	}
	return true, value
}

func (d *DataBase) IsPresentedByURL(url string) (bool, string) {
	var key string
	query := "SELECT key FROM hashTable WHERE value = $1"
	err := d.Db.QueryRow(query, url).Scan(&key)
	if err != nil {
		return false, key
	}
	return true, key
}
