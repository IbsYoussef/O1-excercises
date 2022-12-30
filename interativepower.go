package piscine

func IterativePower(nb int, power int) int {
	multiply := nb
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 20 {
		return 0
	}
	for i := 0; i < power; i++ {
		nb *= multiply
	}
	return nb

}
