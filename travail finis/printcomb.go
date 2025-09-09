package main

import "github.com/01-edu/z01"

func main() {
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {
				z01.PrintRune(numbers[i])
				z01.PrintRune(numbers[j])
				z01.PrintRune(numbers[k])
				if !(i == 7 && j == 8 && k == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
