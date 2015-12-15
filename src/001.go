package main

import "fmt"

func main() {
	ret := []rune("")
	for i, c := range []rune("パタトクカシー") {
		if i%2 == 0 {
			ret = append(ret, c)
		}
	}
	fmt.Println(string(ret))
}
