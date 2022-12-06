package piscine

// package main

import (
	"fmt"
)

func UltimatePointOne() {
	a := 0
	b := &a
	n := &b

	fmt.Println(&n)
	fmt.Println(a)
}
