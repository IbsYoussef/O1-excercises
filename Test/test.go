package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(NRune("Hello!", 5))
}

func NRune(s string, n int) rune {
	if n > len(s) {
		return 0
	}
	Change := []rune(s)
	CharacterToFind := Change[n]
	return CharacterToFind
}
