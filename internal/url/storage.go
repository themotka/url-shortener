package url

import (
	"database/sql"
	_map "themotka/shortener/internal/url/adapters/db/map"
	"themotka/shortener/internal/url/adapters/db/pg"
)

type Storage interface {
	Write(key string, url string) error
	IsPresentedByKey(key string) (bool, string)
	IsPresentedByURL(url string) (bool, string)
}

func NewStorage(usePostgres bool, db *sql.DB) Storage {
	if usePostgres {
		return &pg.DataBase{Db: db}
	} else {
		return &_map.MapUrl{Table: make(map[string]string)}
	}
}
