package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	start := -1

	for i, r := range s {
		if r != ' ' && r != '\t' && r != '\n' {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				words = append(words, s[start:i])
				start = -1
			}
		}
	}

	if start != -1 {
		words = append(words, s[start:])
	}

	return words
}
