package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := 0
	baseFromLen := len(baseFrom)
	for _, r := range nbr {
		for i, br := range baseFrom {
			if r == br {
				num = num*baseFromLen + i
				break
			}
		}
	}

	if num == 0 {
		return string(baseTo[0])
	}

	result := ""
	baseToLen := len(baseTo)
	for num > 0 {
		result = string(baseTo[num%baseToLen]) + result
		num /= baseToLen
	}

	return result
}
