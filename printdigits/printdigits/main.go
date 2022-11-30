package main

import (
	"github.com/01-edu/z01"
	"github.com/01-edu/zo1"
)

func main() {
	for ch := '0'; ch <= '9'; ch++ {
		zo1.PrintRune(ch)
	}

	z01.PrintRune('\n')
}
