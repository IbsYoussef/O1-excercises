package piscine

// package main

import "github.com/01-edu/z01"

func StrLen(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}
