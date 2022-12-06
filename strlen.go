package piscine

import (
	"fmt"
	"unicode/utf8"
)

// package main

// func main (){
// }

func StrLen(s string) int {
	I := ("Hello World")
	fmt.Println(I)
	fmt.Println(utf8.RuneCountInString(s))
}
