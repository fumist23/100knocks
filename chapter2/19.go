package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"sort"
)



type Hightemp struct {
	pref string
	count int
}

type Hightemps []Hightemp

func (t Hightemps) Len() int {
	return len(t)
}

func (t Hightemps) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Hightemps) Less(i, j int) bool{
	return t[i].count > t[j].count
}

func main() {
	filename := "hightemp.txt"
	file, err := os.Open(filename)
	if err !=nil {
		fmt.Printf("os.Open: %#v", err)
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
	words := []string{}
	for _, line := range lines{
		rows := strings.Fields(line)
		word := rows[0]
		words = append(words, word)
	}

	tmp := make(map[string]int)
	for _, word := range words{
		_, exist := tmp[word]
		if exist == false {
			tmp[word] = 1
		}else {
			tmp[word] = tmp[word] + 1
		}
	}

	arr := []Hightemp{}
	for key, value := range tmp{
		t := Hightemp{key, value}
		arr = append(arr, t)
	}

	sort.Sort(Hightemps(arr))
	for _, v := range arr{
		fmt.Printf("%s  %d回\n", v.pref, v.count)
	}
}

