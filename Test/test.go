package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(LastRune("Why"))
}

func LastRune(s string) rune {
	Change := []rune(s)
	LastCharacter := Change[len(s)-1]
	return LastCharacter
}
