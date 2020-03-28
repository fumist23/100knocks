package main 

import (
	"fmt"
)

func main() {
	a := "パトカー"
	b := "タクシー"
	ptc := []rune(a)
	txy := []rune(b)
	size := len(ptc)
	var rn []rune
	for i := 0; i < size; i++ {
		rn = append(rn, ptc[i], txy[i])
	}
	fmt.Printf("%s\n", string(rn))
}