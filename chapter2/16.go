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
		fmt.Printf("strconv.Atoi: %#v", err)
		return
	}

	filename :="hightemp.txt"
	lines := scan(filename)
	
	split := len(lines)/ n  //各ファイルの行数

	newlines := []string{}
	for i := 1; i <= len(lines); i++ {
		newlines = append(newlines, lines[i-1])
		if i%split == 0 {
			write(newlines, strconv.Itoa(i) + ".txt")
			newlines = []string{}
		}
	}
}

func scan(filename string)([]string){
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v", err)
	}
	return lines
}

func write(row []string, newFile string){
	f, err := os.Create(newFile)
	if err != nil {
		fmt.Printf("os.Create: %#v", err)
		return
	}
	writer := bufio.NewWriter(f)
	for _, v := range row{
		fmt.Fprint(writer, v, "\n")
	}
	writer.Flush()
	return
}
