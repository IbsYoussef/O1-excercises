package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(NRune("Hello!", 4))
}

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		Change := []rune(s)
		CharacterToFind := Change[n-1]
		return CharacterToFind
	}
	return 0
}
