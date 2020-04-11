package main 

import(
	"encoding/json"
	"os"
	"bufio"
	"io"
	"fmt"
)

type Article struct{
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
	for _, article := range articles{
		if article.Title == "イギリス" {
			fmt.Println(article.Text)
		}
	}
}