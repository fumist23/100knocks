package main 

import "fmt"

func main() {
	s := "パタトクカシーー"
	rn := []rune(s)
	var even []rune
	for i, v := range rn{
		if i%2 == 0 {
			even = append(even, v)
		}
	}
	fmt.Printf("%s\n", string(even))
}
