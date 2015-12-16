package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	words := strings.Split(s, " ")
	ret := []int{}
	for _, word := range words {
		ret = append(ret, len(strings.Trim(word, ",.")))
	}
	fmt.Println(ret)
}
