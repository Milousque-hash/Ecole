package piscine

func Index(s, substr string) int {
	rs := []rune(s)
	rsub := []rune(substr)
	if len(rsub) == 0 {
		return 0
	}
	for i := 0; i <= len(rs)-len(rsub); i++ {
		match := true
		for j := 0; j < len(rsub); j++ {
			if rs[i+j] != rsub[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
