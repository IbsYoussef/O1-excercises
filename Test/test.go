package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	fmt.Println(StrRev(s))
}

func StrRev(s string) string {
	var Reverse string
	for i := len(s) - 1; i > 0; i-- {
		Reverse += string(s[i])
	}
	return Reverse
}
