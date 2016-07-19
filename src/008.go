package main

import (
	"fmt"
	"unicode"
)

func cipher(s string) string {
	var chars []rune
	for _, char := range s {
		if unicode.IsLower(char) {
			chars = append(chars, rune(219 - char))
		} else {
			chars = append(chars, char)
		}
	}
	return string(chars)
}

func main() {
	const s = "This is a TEST sentence."
	fmt.Println(s)
	fmt.Println(cipher(s))
	fmt.Println(cipher(cipher(s)))
}
