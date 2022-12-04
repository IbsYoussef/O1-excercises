package piscine

// package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i < j && j < k && !(i == 7 && j == 8 && k == 9) {
					z01.PrintRune('1')
					z01.PrintRune('2')
					z01.PrintRune('3')
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else if i == 7 && j == 8 && k == 9 {
					z01.PrintRune('4')
					z01.PrintRune('5')
					z01.PrintRune('6')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
