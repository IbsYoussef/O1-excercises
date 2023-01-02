package piscine

/* func LastRune(s string) rune {
	Change := []rune(s)
	var Counter rune = '0'
	for i := 0; i <= len(s); i++ {
		Counter++
	}
	Change[0] = rune(Counter)
	return Counter
} */
// My Original Method

func LastRune(s string) rune {
	Change := []rune(s)
	LastCharacter := Change[len(s)-1]
	return LastCharacter
}
