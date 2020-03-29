package main

import (
	"fmt"
)

func main(){
	x := 12
	y := "気温"
	var z float32 
	z = 22.4
	fmt.Println(temperature(x, y, z))
}

func temperature(x int, y string, z float32)(string){
	temperature := fmt.Sprintf("%d時の%sは%vv", x, y, z)
	return temperature
}