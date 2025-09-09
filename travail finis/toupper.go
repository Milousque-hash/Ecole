package piscine

func ToUpper(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r - ('a' - 'A')
		}
	}
	return string(runes)
}
