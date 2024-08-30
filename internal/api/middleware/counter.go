package middleware

import (
	"database/sql"
	"themotka/shortener/internal/database"
)

type HashTable struct {
	Table       map[string]string
	Repo        *database.Repository
	charSet     []rune
	tableLength int
	keyLength   int
	keys        []string
}

func NewHashTable(db *sql.DB) *HashTable {
	h := &HashTable{Table: make(map[string]string),
		Repo:        database.NewRepository(db),
		charSet:     []rune("0123456789abcdefghijklmnopqrstuvwxyz"),
		tableLength: 0,
		keyLength:   1,
		keys:        make([]string, 0)}
	h.generateKeys(1)
	return h
}

func (h *HashTable) WriteTo(url string) string {
	if h.Repo == nil {
		for k, v := range h.Table {
			if v == url {
				return k
			}
		}
	} else {
		if k, _ := h.Repo.GetKeyByValue(url); k != "" {
			return k
		}
	}

	h.tableLength++
	if pow(36, h.keyLength) <= h.tableLength {
		h.keyLength++
		h.generateKeys(h.keyLength)
	}
	key := h.keys[h.tableLength-1]
	if h.Repo == nil {
		h.Table[key] = url
	} else {
		err := h.Repo.Put(key, url)
		if err != nil {
			return err.Error()
		}
	}
	return key
}

func (h *HashTable) generateKeys(length int) {
	n := len(h.charSet)
	for i := 0; i < pow(n, length); i++ {
		current := ""
		num := i
		for j := 0; j < length; j++ {
			current = string(h.charSet[num%n]) + current
			num /= n
		}
		h.keys = append(h.keys, current)
	}
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
