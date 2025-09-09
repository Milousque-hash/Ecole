package piscine

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	i := 0
	for i < len(ra) && i < len(rb) {
		if ra[i] > rb[i] {
			return 1
		} else if ra[i] < rb[i] {
			return -1
		}
		i++
	}
	if len(ra) > len(rb) {
		return 1
	} else if len(ra) < len(rb) {
		return -1
	}
	return 0
}
