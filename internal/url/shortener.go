package url

type Shortener interface {
	Shorten() string
}

type shortener struct {
	charSet     []rune
	tableLength int
	keyLength   int
	keys        []string
}

func NewShortener() Shortener {
	h := &shortener{
		charSet:     []rune("0123456789abcdefghijklmnopqrstuvwxyz"),
		tableLength: 0,
		keyLength:   1,
		keys:        make([]string, 0)}
	h.generateKeys(1)
	return h
}

func (s *shortener) Shorten() string {
	s.tableLength++
	if pow(36, s.keyLength) <= s.tableLength {
		s.keyLength++
		s.generateKeys(s.keyLength)
	}
	key := s.keys[s.tableLength-1]
	return key
}

func (s *shortener) generateKeys(length int) {
	n := len(s.charSet)
	for i := 0; i < pow(n, length); i++ {
		current := ""
		num := i
		for j := 0; j < length; j++ {
			current = string(s.charSet[num%n]) + current
			num /= n
		}
		s.keys = append(s.keys, current)
	}
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
