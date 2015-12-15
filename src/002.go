package main

import "fmt"

func main() {
	s1 := []rune("パトカー")
	s2 := []rune("タクシー")
	ret := []rune{}
	for i := 0; i < len(s1); i++ {
		ret = append(ret, s1[i], s2[i])
	}
	fmt.Println(string(ret))
}
