package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	col1 := "col1.txt"
	col2 := "col2.txt"
	col := "col.txt"

	c1, err := os.Open(col1)
	if err != nil {
		fmt.Printf("os.Open1: %#v", err)
		return
	}
	defer c1.Close()

	c2, err := os.Open(col2)
	if err != nil {
		fmt.Printf("os.Open2: %#v", err)
		return
	}
	defer c2.Close()

	c, err := os.Create(col)
	if err != nil {
		fmt.Printf("os.Create: %#v", err)
		return
	}
	defer c.Close()

	scanner1 := bufio.NewScanner(c1)
	scanner2 := bufio.NewScanner(c2)

	for scanner1.Scan() {
		scanner2.Scan()
		c.WriteString(scanner1.Text() + "\t" + scanner2.Text() + "\n")
	}

	if err := scanner1.Err(); err != nil {
		fmt.Printf("scanner1.Err: %#v", err)
		return
	}

	if err := scanner2.Err(); err != nil {
		fmt.Printf("scanner2.Err: %#v", err)
		return
	}
}