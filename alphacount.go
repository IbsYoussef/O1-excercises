package piscine

func AlphaCount(s string) int {
	Slice := []rune(s)
	Counter := 0
	for i := 0; i < len(s); i++ {
		if Slice[i] > 64 && Slice[i] < 91 || Slice[i] > 96 && Slice[i] < 123 {
			Counter += 1
		}
	}
	return Counter
}
