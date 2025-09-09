package main

import "github.com/01-edu/z01"

func main() {
	numbers := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i1 := 0; i1 < 10; i1++ {
		for i2 := 0; i2 < 10; i2++ {
			first := i1*10 + i2
			for j1 := 0; j1 < 10; j1++ {
				for j2 := 0; j2 < 10; j2++ {
					second := j1*10 + j2
					if first < second {
						z01.PrintRune(numbers[i1])
						z01.PrintRune(numbers[i2])
						z01.PrintRune(' ')
						z01.PrintRune(numbers[j1])
						z01.PrintRune(numbers[j2])
						if !(first == 98 && second == 99) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
