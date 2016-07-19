package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	words := strings.Split(s, " ")
	indexes := map[int]int{1:0, 5:0, 6:0, 7:0, 8:0, 9:0, 15:0, 16:0, 19:0}
	ret := map[string]int{}
	for i, word := range words {
		if _, ok := indexes[i+1]; ok {
			ret[word[:1]] = i + 1
		} else {
			ret[word[:2]] = i + 1
		}
	}
	fmt.Println(ret)
}
