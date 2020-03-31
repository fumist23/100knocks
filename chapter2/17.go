package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	filename :="hightemp.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner.Err: %#v", err)
		return
	}

	sort := make(map[string]bool)
	for _, v := range lines{
		rows := strings.Fields(v)
		_, exist := sort[rows[0]]
		if exist == false {
			sort[rows[0]] = true
		}
	}
	for key, _ := range sort{
		fmt.Println(key)
	}
}