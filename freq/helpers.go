package freq

import "strings"

func countLetters(b []byte) int {
	s := strings.ToUpper(string(b))
	count := 0
	for i := range s {
		if strings.ContainsRune(string(Alphabet), rune(s[i])) {
			count++
		}
	}
	return count
}

func countSpaces(b []byte) int {
	count := 0
	s := string(b)
	for i := range s {
		if s[i] == ' ' {
			count++
		}
	}
	return count
}

func letterProb(l byte, b []byte) float64 {
	count := 0
	letters := countLetters(b)
	s := strings.ToUpper(string(b))
	for i := 0; i < len(s); i++ {
		if s[i] == l {
			count++
		}
	}
	return float64(count) / float64(letters)
}

func CL(b []byte) int {
	s := strings.ToUpper(string(b))
	count := 0
	for i := range s {
		if strings.ContainsRune(string(Alphabet), rune(s[i])) {
			count++
		}
	}
	return count
}

func CS(b []byte) int {
	count := 0
	s := string(b)
	for i := range s {
		if s[i] == ' ' {
			count++
		}
	}
	return count
}

func CO(b []byte) int {
	let := CL(b)
	sp := CS(b)
	return len(b) - let - sp
}
