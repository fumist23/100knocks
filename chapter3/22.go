package main 

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"io"
	"encoding/json"
)

	type Article struct {
		Text string `json:"text"`
		Title string `json:"title"`
	}


func main() {
	filename := "jawiki-country.json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	articles := []Article{}
	for {
		b, err := reader.ReadBytes('\n')
		if err == io.EOF {
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

	reg, _ := regexp.Compile(`Category.*]`)
	for _, v := range reg.FindAll([]byte(txt), -1){
		fmt.Println(string(v[len("Category:"):len(v)-2]))
	}
}