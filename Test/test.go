package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(FirstRune("Hello"))
}

func FirstRune(s string) rune {
	FirstRuneLetter := []rune(s)
	result := FirstRuneLetter[0]
	return result
}
