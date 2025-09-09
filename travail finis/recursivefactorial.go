package piscine

import "github.com/01-edu/z01"

func PrintCombn(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	comb := make([]rune, n)

	var generate func(pos, start int)
	generate = func(pos, start int) {
		if pos == n {
			for i := 0; i < n; i++ {
				z01.PrintRune(comb[i])
			}
			// Vérifie si c'est la dernière combinaison
			last := true
			for i := 0; i < n; i++ {
				if comb[i] != numbers[10-n+i] {
					last = false
					break
				}
			}
			if !last {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}
		for i := start; i < 10; i++ {
			comb[pos] = numbers[i]
			generate(pos+1, i+1)
		}
	}

	generate(0, 0)
	z01.PrintRune('\n')
}
