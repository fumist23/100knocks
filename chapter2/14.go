package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)



func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("strconv.Atoi(os.Args[1]): %#v", err)
		return
	}

	filename := "hightemp.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; i < n+1; i++ {
		scanner.Scan()
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v", err)
		return
	}
}