package main 

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main() {
	randInit()

	sentence := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

	fmt.Println(typoglycemia(sentence))
}

func typoglycemia(str string) string{
	words := strings.Split(str, " ")
	var typo string
	for _, word := range words{
		if len(word) > 4 {
			rn := []rune(word)
			shaffle(rn[1:len(word)-1])
			word = string(rn)
		}

		typo += word + " "
	}
	return typo
}


func randInit(){
	rand.Seed(time.Now().UnixNano())
}

func shaffle(rn []rune)[]rune{
	n := len(rn)
	for i := n-1; 0 < i; i-- {
		j := rand.Intn(i + 1)
		rn[i], rn[j] = rn[j], rn[i]
	}
	return rn 
}