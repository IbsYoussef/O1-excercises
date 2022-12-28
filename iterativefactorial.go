package piscine

// package main

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := 0; i <= nb; i++ {
		result *= i
	}
	return result
}
