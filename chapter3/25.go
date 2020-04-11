package main 

import (
	"os"
	"io"
	"encoding/json"
	"fmt"
	"regexp"
	"bufio"
	"strings"
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

	reg, _ := regexp.Compile(`{{基礎情報 国[\s\S]*\n}}`)
	txt = string(reg.FindAll([]byte(txt), -1)[0])
	txt = strings.Replace(txt, "{{基礎情報 国\n", "", -1)
	txt = strings.Replace(txt, "\n|", "\n|\n|", -1)
	txt = strings.Replace(txt, "\n}}", "\n|}}", -1)

	fields := make(map[string]string)
	reg, err = regexp.Compile(`(?m)^\|[\s\S]*?\n\|`)
	for _, v := range reg.FindAll([]byte(txt), -1){
		s := string(v[1:len(v)-2])
		two := strings.Split(s, " = ")
		fields[two[0]] = two[1]
	}

}	