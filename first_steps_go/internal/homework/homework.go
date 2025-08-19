package homework

import (
	"errors"
	"unicode"
)

func checkStringIsSpaces(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func checkStringIsEmpty(s string) bool {
	return len(s) == 0
}

func WordCount(s string) (map[string]int, error) {
	if checkStringIsEmpty(s) {
		return nil, errors.New("empty string")
	}

	if checkStringIsSpaces(s) {
		return nil, errors.New("string is only spaces")
	}

	res := make(map[string]int)
	start := -1
	for i, r := range s {
		if unicode.IsSpace(r) {
			if start >= 0 {
				res[s[start:i]]++
				start = -1
			}
			continue
		}
		if start < 0 {
			start = i
		}
	}
	if start >= 0 {
		res[s[start:]]++
	}
	return res, nil
}

func AreAnagrams(s1, s2 string) (bool, error) {
	if checkStringIsEmpty(s1) || checkStringIsEmpty(s2) {
		return false, errors.New("empty string")
	}

	if checkStringIsSpaces(s1) || checkStringIsSpaces(s2) {
		return false, errors.New("string is only spaces")
	}

	count := make(map[rune]int)

	len1, len2 := 0, 0

	for _, r := range s1 {
		if unicode.IsSpace(r) {
			continue
		}
		count[unicode.ToLower(r)]++
		len1++
	}

	for _, r := range s2 {
		if unicode.IsSpace(r) {
			continue
		}
		cr := unicode.ToLower(r)
		count[cr]--
		if count[cr] < 0 {
			return false, nil
		}
		len2++
	}

	return len1 == len2, nil
}

func FirstUnique(s string) (string, error) {
	if checkStringIsEmpty(s) {
		return "0", errors.New("empty string")
	}

	if checkStringIsSpaces(s) {
		return "0", errors.New("string is only spaces")
	}

	count := make(map[rune]int, len(s))
	for _, r := range s {
		count[r]++
	}

	for _, r := range s {
		if count[r] == 1 {
			return string(r), nil
		}
	}
	return "0", nil
}
