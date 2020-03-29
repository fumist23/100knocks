package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	filename := "hightemp.txt"
	col1 := "col1.txt"
	col2 := "col2.txt"
	
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	c1, err := os.Create(col1)
	if err != nil {
		fmt.Printf("os.create: %#v", err)
		return
	}
	defer c1.Close()

	c2, err := os.Create(col2)
	if err != nil {
		fmt.Printf("os.create: %#v", err)
	}
	defer c2.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		separate := strings.Split(scanner.Text(), "\t")
		c1.WriteString(separate[0] + "\n")
		c2.WriteString(separate[1] + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v", err)
		return
	}
}