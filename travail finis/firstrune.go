package piscine

func FirstRune(str string) rune {
	for _, r := range str {
		return r
	}
	return 0
}
