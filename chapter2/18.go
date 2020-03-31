package main 

import (
	"fmt"
	"sort"
	"os"
	"strconv"
	"strings"
	"bufio"
)


type Hightemp struct {
	pref string
	city string
	temp float64
	time string
}

type Hightemps []Hightemp

func (t Hightemps) Len() int {
	return len(t)
}

func (t Hightemps) Swap(i,j int){
	t[i], t[j] = t[j], t[i]
}

func (t Hightemps) Less(i, j int)bool {
	return t[i].temp > t[j].temp
}

func main(){
	filename := "hightemp.txt"
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
	if err := scanner.Err(); err != nil{
		fmt.Printf("scanner.Err: %#v", err)
		return
	}

	arr := []Hightemp{}
	for _, line := range lines{
		rows := strings.Fields(line)
		change, _ := strconv.ParseFloat(rows[2], 64)
		tmp := Hightemp{rows[0], rows[1], change, rows[3]}
		arr = append(arr, tmp)
	}
	sort.Sort(Hightemps(arr))
	for _, line := range arr{
		fmt.Printf("%s\t%s\t%v\t%s\n", line.pref, line.city, line.temp, line.time)
	}
}