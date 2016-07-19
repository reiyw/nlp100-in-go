package main

import (
	"fmt"
	"strings"
)

func wordNGram(str string, n int) [][]string {
	words := strings.Split(str, " ")

	extended := []string{}
	for i := 0; i < n - 1; i++ {
		extended = append(extended, "")
	}
	extended = append(extended, words...)
	for i := 0; i < n - 1; i++ {
		extended = append(extended, "")
	}

	var ret [][]string
	for i := 0; i < len(words) + n - 1; i++ {
		ret = append(ret, extended[i:i+n])
	}

	return ret
}

func charNGram(str string, n int) []string {
	chars := strings.Repeat(" ", n - 1) + str + strings.Repeat(" ", n - 1)

	var ret []string
	for i := 0; i < len(str) + n - 1; i++ {
		ret = append(ret, chars[i:i+n])
	}

	return ret
}

func main() {
	const s = "I am an NLPer"
	fmt.Println(wordNGram(s, 2))
	fmt.Println(charNGram(s, 2))
}
