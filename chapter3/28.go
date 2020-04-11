package main 

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"encoding/json"
	"strings"
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

	reg := regexp.MustCompile(`{{基礎情報 国[\s\S]*\n}}`)
	txt = string(reg.FindAll([]byte(txt), -1)[0])
	txt = strings.Replace(txt, "{{基礎情報 国\n", "", -1)
	txt = strings.Replace(txt, "\n|", "\n|\n|", -1)
	txt = strings.Replace(txt, "\n}}", "\n|}}", -1)
	reg = regexp.MustCompile(`(?m)^\|[\s\S]*?\n\|`)
	regEmphasis := regexp.MustCompile(`'{2,5}`)
	regInternalLink := regexp.MustCompile(`\[\[.*?\]\]`)
	regInternalLink1 := regexp.MustCompile(`(.*)\[\[(.*)#.*\]\](.*)`)
	regInternalLink2 := regexp.MustCompile(`(.*)\[\[(.*)\|.*\]\](.*)`)
	regExternalLink := regexp.MustCompile(`\[?http.*\]?`)
	regCommentOut := regexp.MustCompile(`<!--.*-->`)
	fields := make(map[string]string)
	for _, v := range reg.FindAll([]byte(txt), -1){
		eachline := string(v[1:len(v)-2])
		elements := strings.Split(eachline, " = ")
		removedEmphasis := regEmphasis.ReplaceAllString(elements[1], "")

		if regInternalLink.FindString(removedEmphasis) == ""{
			fields[elements[0]] = removedEmphasis
		}else if regInternalLink1.FindString(removedEmphasis) != "" {
			for _, s := range regInternalLink1.FindAllStringSubmatch(removedEmphasis, -1){
				var str string
				for i := 1; i < len(s); i++{
					str += s[i] 
				}
				str = strings.Replace(str, "[[", "", -1)
				str = strings.Replace(str, "[[", "", -1)
				fields[elements[0]] = str
			}
		}else if regInternalLink2.FindString(removedEmphasis) != ""{
			for _, s := range regInternalLink2.FindAllStringSubmatch(removedEmphasis, -1){
				var str string
				for i := 1; i < len(s); i++{
					str += s[i] 
				}
				str = strings.Replace(str, "[[", "", -1)
				str = strings.Replace(str, "]]", "", -1)
				fields[elements[0]] = str
			}
		}else{
			s := removedEmphasis
			s = strings.Replace(s, "[[", "", -1)
			s = strings.Replace(s, "]]", "", -1)
			fields[elements[0]] = s
		}

		fields[elements[0]] = regExternalLink.ReplaceAllString(fields[elements[0]], "")
		fields[elements[0]] = regCommentOut.ReplaceAllString(fields[elements[0]], "")
	}

	for _, value := range fields{
		fmt.Println(value)
	}
	
}