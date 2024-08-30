package middleware

import (
	"fmt"
)

type HashTable struct {
	table       map[string]string
	charSet     []rune
	tableLength int
	keys        []string
}

func NewHashTable() *HashTable {
	return &HashTable{
		table:       make(map[string]string),
		charSet:     []rune("0123456789abcdefghijklmnopqrstuvwxyz"),
		tableLength: 1,
		keys:        make([]string, 0)}
}

func (h *HashTable) WriteTo(url string) error {

	h.generateKeys(1)
	h.generateKeys(2)
	for _, combination := range h.keys {
		fmt.Println(combination)
	}
	return nil
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
