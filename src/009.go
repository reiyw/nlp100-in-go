package main

import (
	"fmt"
	"strings"
	"math/rand"
)

func typoglycemia(s string) string {
	words := strings.Split(s, " ")
	var retWords []string
	for _, word := range words {
		if len(word) > 4 {
			n := len(word)
			dest := make([]byte, n)
			perm := rand.Perm(len(word) - 2)
			dest[0] = word[0]
			for i, v := range perm {
				dest[v + 1] = word[i + 1]
			}
			dest[n - 1] = word[n - 1]
			retWords = append(retWords, string(dest))
		} else {
			retWords = append(retWords, word)
		}
	}
	return strings.Join(retWords, " ")
}

func main() {
	const s = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	fmt.Println(typoglycemia(s))
}
