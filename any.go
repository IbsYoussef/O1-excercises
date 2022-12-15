package piscine

// func main (){}

func Any(f func(string) bool, a []string) bool {
	for _, element := range a {
		resultat := f(element)
		if resultat {
			return true
		}
	}
	return false
}
