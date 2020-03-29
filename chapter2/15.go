package main 

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("strconv.Atoi(os.Args[0]): %#v", err)
		return
	}

	filename := "hightemp.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, 24)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	for i := 1; i < n+1; i++ {
		fmt.Println(lines[len(lines)-n-1+i])
	}
	
}

