package main

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb < 21 {
		var result int = 1
		for num := 1; num <= nb; num++ {
			result = num * result
		}
		return result
	} else {
		return 0
	}
}
