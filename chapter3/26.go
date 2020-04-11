package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"regexp"
	"io"
	"encoding/json"
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

	reg := regexp.MustCompile(`{{基礎情報 国[\s\S]*\n}}`)
	txt = string(reg.FindAll([]byte(txt), -1)[0])
	txt = strings.Replace(txt, "{{基礎情報 国\n", "", -1)
	txt = strings.Replace(txt, "\n|", "\n|\n|", -1)
	txt = strings.Replace(txt, "\n}}", "\n|}}", -1)

	reg = regexp.MustCompile(`(?m)^\|[\s\S]*?\n\|`)
	fields := make(map[string]string)
	regEmphasis := regexp.MustCompile(`'{2,5}`)
	for _, v := range reg.FindAll([]byte(txt), -1){
		eachline := string(v[1:len(v)-2])
		elements := strings.Split(eachline, " = ")
		fields[elements[0]] = regEmphasis.ReplaceAllString(elements[1], "")
	}
	fmt.Println(fields["確立形態4"])  //確認用
}