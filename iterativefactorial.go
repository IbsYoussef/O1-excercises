package piscine

func IterativeFactorial(nb int) int {
	compteur := 1
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		compteur *= i
	}
	return compteur
}
