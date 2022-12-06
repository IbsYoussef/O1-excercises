package piscine

import (
	"unicode/utf8"
	"fmt"
	"piscine"
)

// package main

// func main (){
// }

func (s string) int{
	s := piscine.Strlen("Hello World")
	fmt.Println(s, utf8.RuneCountInString(str))
}