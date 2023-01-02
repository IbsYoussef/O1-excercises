package piscine

func NRune(s string, n int) rune {
	if n > len(s) && n < 0 {
		return 0
	}
	Change := []rune(s)
	CharacterToFind := Change[n-1]
	return CharacterToFind
}
