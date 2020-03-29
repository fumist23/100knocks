package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	filename := "hightemp.txt"
	line := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v\n" ,err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line++ 
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v\n", err)
		return
	}

	fmt.Println(line)
}