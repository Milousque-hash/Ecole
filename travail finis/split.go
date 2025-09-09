package piscine

func Split(s string, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	start := 0
	sepLen := len(sep)

	for i := 0; i <= len(s)-sepLen; {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}

	result = append(result, s[start:])

	return result
}
