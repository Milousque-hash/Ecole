package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	foundFirstDigit := false
	for _, r := range s {
		if !foundFirstDigit {
			if r == '-' {
				sign = -1
			} else if r >= '0' && r <= '9' {
				foundFirstDigit = true
				result = int(r - '0')
			}
		} else {
			if r >= '0' && r <= '9' {
				result = result*10 + int(r-'0')
			}
		}
	}
	return result * sign
}
