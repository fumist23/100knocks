package main 

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	sentence = strings.Replace(sentence, ",", "", -1)
	sentence = strings.Trim(sentence, ".")
	words := strings.Fields(sentence)

	for i := 0; i < len(words); i++ {
		fmt.Printf("%v ", len(words[i]))
	}
	fmt.Printf("\n")
}