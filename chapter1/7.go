package main 

import (
	"fmt"
	
)

func main(){
	x := 12
	y := "気温"
	z := 22.4
	t := temperature(x, y, z)
	fmt.Println(t)
}

func temperature(x int, y string, z float64){
	 fmt.Printf("%d 時の %s は %f", x, y, z)
	 return
}