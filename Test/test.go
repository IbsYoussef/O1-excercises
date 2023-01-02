package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(NRune("Hello!", 10))
}

func NRune(s string, n int) rune {
	if n > len(s) || n < 0 {
		return 0
	}
	Change := []rune(s)
	NCharacterToFind := Change[n-1]
	return NCharacterToFind
}
