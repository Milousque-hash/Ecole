package piscine

func NRune(str string, n int) rune {
	i := 1
	for _, r := range str {
		if i == n {
			return r
		}
		i++
	}
	return 0
}
