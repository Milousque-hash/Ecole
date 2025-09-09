package piscine

func IsLower(str string) bool {
	if str == "" {
		return false
	}
	for _, r := range str {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
