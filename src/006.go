package main

import (
	"fmt"
	"strings"
	"github.com/emirpasic/gods/sets/hashset"
)

func charNGram(str string, n int) []string {
	chars := strings.Repeat(" ", n - 1) + str + strings.Repeat(" ", n - 1)

	var ret []string
	for i := 0; i < len(str) + n - 1; i++ {
		ret = append(ret, chars[i:i+n])
	}

	return ret
}

func main() {
	const s1 = "paraparaparadise"
	const s2 = "paragraph"

	setX := hashset.New()
	for _, char := range charNGram(s1, 2) {
		setX.Add(char)
	}

	setY := hashset.New()
	for _, char := range charNGram(s2, 2) {
		setY.Add(char)
	}

	fmt.Println(setX)
	fmt.Println(setY)

	union := hashset.New()
	for _, char := range charNGram(s1, 2) {
		union.Add(char)
	}
	for _, char := range charNGram(s2, 2) {
		union.Add(char)
	}

	fmt.Println(union)

	intersection := hashset.New()
	for _, char1 := range charNGram(s1, 2) {
		for _, char2 := range charNGram(s2, 2) {
			if char1 == char2 {
				intersection.Add(char1)
			}
		}
	}

	fmt.Println(intersection)

	fmt.Println(setX.Contains("se"))
	fmt.Println(setY.Contains("se"))

	for _, char := range intersection.Values() {
		setX.Remove(char)
		setY.Remove(char)
	}

	fmt.Println(setX)
	fmt.Println(setY)
}
