package main 

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)
func main(){
	filename := "hightemp.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.Replace(scanner.Text(), "\t", " ", -1))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v", err)
		return
	}

}
