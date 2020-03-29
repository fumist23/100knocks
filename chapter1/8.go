package main

import (
	"fmt"
)

func main(){
	sentence := "Never let your memories be greater than your dreams"
	crypted := cipher(sentence)
	fmt.Println(crypted)
}

func cipher(s string)string {
	var x rune
	var rn []rune
	for _, v := range s{
		if v >= 97 && v <= 122 {
			x = 219 - v
		}else {
			x = v
		}
		rn = append(rn, x)
	}
	return string(rn)
}
