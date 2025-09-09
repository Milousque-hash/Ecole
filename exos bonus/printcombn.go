package piscine

import "github.com/01-edu/z01"

func PrintCombNRecursive(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	comb := make([]int, n)
	generate(0, 0, n, comb)
	z01.PrintRune('\n')
}

func generate(pos, start, n int, comb []int) {
	if pos == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune('0' + comb[i]))
		}
		if comb[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	for i := start; i <= 10-n+pos; i++ {
		comb[pos] = i
		generate(pos+1, i+1, n, comb)
	}
}
