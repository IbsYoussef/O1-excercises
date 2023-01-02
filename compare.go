package piscine

func Compare(a, b string) int {
	Astring := len(a)
	Bstring := len(b)
	var tour int
	if Astring > Bstring {
		tour = Bstring
	} else {
		tour = Astring
	}
	for i := 0; i < tour; i++ {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	return 0
}
