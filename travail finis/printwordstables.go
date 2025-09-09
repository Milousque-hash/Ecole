package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(words []string) {
	for _, word := range words {
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
