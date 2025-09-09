package piscine

import "github.com/01-edu/z01"

func PrintNbr(nbr int) {
	if nbr == 0 {
		z01.PrintRune('0')
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
	}

	num := nbr
	digits := []int{}

	for num != 0 {
		last := num % 10
		if last < 0 {
			last = -last
		}
		digits = append(digits, last)
		num = num / 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
