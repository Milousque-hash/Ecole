package piscine

import "github.com/01-edu/z01"

func LastRune(str string) rune {
	var last rune
	for _, r := range str {
		last = r
	}
	return last
}
