package main 

import (
	"fmt"
	"strings"
)

func main(){
	a := "paraparaparadise"
	b := "paragraph"
	a2 := strings.Split(a, "")
	fmt.Println(a2)
	b2 := strings.Split(b, "")
	fmt.Println(b2)
	b3 := strings.Join(b2, "")
	fmt.Println(b3)
}
