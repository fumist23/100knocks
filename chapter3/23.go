package main 

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"encoding/json"
	"regexp"
)

type Article struct {
	Text string `json:"text"`
	Title string `json:"title"`
}

func main(){
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

	reg, _ := regexp.Compile(`(?m)^=+.*=+`)
	for _, v := range reg.FindAll([]byte(txt), -1){
																  //ここからの処理がわかんないよおおおおおおおお
																
	}
}