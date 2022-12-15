package piscine

// func main(){}

func CountIf(f func(string) bool, a []string) int {
	count := 0
	for _, element := range a {
		resultat := f(element)
		if resultat {
			count += 1
		}
	}
	return count
}
