package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host string
	Port string
	User string
	Pass string
	Name string
	Mode string
}

func NewTable(mt *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", mt.Host, mt.Port, mt.User, mt.Pass, mt.Name, mt.Mode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db, nil
}
