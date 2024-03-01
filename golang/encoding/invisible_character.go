package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "‪0312345678"

	fmt.Println(str)

	for i, rune := range str {
		if !unicode.IsPrint(rune) {
			fmt.Printf("Non-printable character at position %d: %U\n", i, rune)
		}
	}
}
