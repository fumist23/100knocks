package main 

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	sentence = strings.Replace(sentence, ".", "", -1)
	words := strings.Fields(sentence)

	elements := make(map[string]int)
	for i := 0; i < len(words); i++ {
		words_rune := []rune(words[i])
		switch i+1 {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19:
			elements[string(words_rune[0])] = i
		default:
			elements[string(words_rune[0:2])] = i
		}
	}
	fmt.Println(elements)
}