package piscine

func Sqrt(nb int) int {
	for i := nb; i > 0; i-- {
		root := i * i
		if root == nb {
			return i
		}
	}
	return 0
}
