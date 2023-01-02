package piscine

func NRune(s string, n int) rune {
	if n > len(s) {
		return 0
	}
	Change := []rune(s)
	CharacterToFind := Change[n]
	return CharacterToFind
}
