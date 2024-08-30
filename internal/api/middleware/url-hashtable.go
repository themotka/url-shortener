package middleware

type HashTable struct {
	Table       map[string]string
	charSet     []rune
	tableLength int
	keyLength   int
	keys        []string
}

func NewHashTable() *HashTable {
	h := &HashTable{Table: make(map[string]string),
		charSet:     []rune("0123456789abcdefghijklmnopqrstuvwxyz"),
		tableLength: 0,
		keyLength:   1,
		keys:        make([]string, 0)}
	h.generateKeys(1)
	return h
}

func (h *HashTable) WriteTo(url string) string {
	for k, v := range h.Table {
		if v == url {
			return k
		}
	}
	h.tableLength++
	if pow(36, h.keyLength) <= h.tableLength {
		h.keyLength++
		h.generateKeys(h.keyLength)
	}
	key := h.keys[h.tableLength-1]
	h.Table[key] = url
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
