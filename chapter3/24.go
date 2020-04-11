package main 

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	"regexp"
	"io"

)

type Article struct{
	Text string `json:"text"`
	Title string `json:"title"`
}

func main() {
	filename := "jawiki-country.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	articles := []Article{}
	for {
		b, err := reader.ReadBytes('\n')
		if err == io.EOF{
			break
		}
		a := Article{}
		json.Unmarshal([]byte(b), &a)
		articles = append(articles, a)
	}

	var txt string
	for _, article := range articles{
		if article.Title == "イギリス"{
			txt = article.Text
		}
	}

	reg, _ := regexp.Compile(`(File|ファイル):.*?\|`)
	for _, v := range reg.FindAll([]byte(txt), -1){
		if string(v[0]) == "F" {
			fmt.Println(string(v[len("File:"):len(v)-1]))
		}else{
			fmt.Println(string(v[len("ファイル:"):len(v)-1]))
		}
	}
}