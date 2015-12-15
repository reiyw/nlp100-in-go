package main

import "fmt"

func main() {
	chars := []rune("stressed")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	fmt.Println(string(chars))
}
