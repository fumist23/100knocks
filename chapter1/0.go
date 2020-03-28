package main 

import "fmt"

func main() {
	s := "stressed"
	rn := []rune(s)
	size := len(rn)

	for a, b := 0, size-1 ; a < b ; a, b = a+1, b-1 {
		rn[a], rn[b] = rn[b], rn[a]
	}
	fmt.Printf("%s\n", string(rn))
}
