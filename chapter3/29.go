package main 

import (
	"fmt"
	"io"
	"os"
	"bufio"
	"encoding/json"
	"strings"
	"regexp"
	"net/url"
	"net/http"
	"log"
	"io/ioutil"
)

type Article struct {
	Text string `json:"text"`
	Title string `json:"title"`
}

type Client struct{
	URL *url.URL
}

func NewClient(urlstr string)(*Client, error){
	parsedURL, err := url.ParseRequestURI(urlstr)
	if err != nil{
		panic(err)
	}

	return &Client{parsedURL}, nil
}

type ResponceWiki struct{
	Batchcomplete string `json:"batchcomplete"`
	Query struct{
		Pages struct{
			Num23473560 struct {
				Pageid int `json:"pageid"`
				Ns int `json:"ns"`
				Title string `json:"title"`
				Imagerepository string `json:"imagerepository"`
				Imageinfo []struct{
					URL string `json:"url"`
					Descriptionurl string `json:"description"`
					Descriptionshorturl string `json:"descriptionshorturl"`
				} `json:"imageinfo"` 
			}`json:"23473560"`
		}`json:"pages"`
	}`json:"query"`
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

	m := make(map[string]string)
	regEmphasis := regexp.MustCompile(`'{2,5}`)
	regInternalLink := regexp.MustCompile(`\[\[.*?\]\]`)
	regInternalLink1 := regexp.MustCompile(`(.*)\[\[(.*)#.*\]\](.*)`)
	regInternalLink2 := regexp.MustCompile(`(.*)\[\[(.*)\|.*\]\](.*)`)
	regExternalLink := regexp.MustCompile(`\[?http.*]?`)
	regCommentOut := regexp.MustCompile(`<!--.*-->`)
	for _, v := range reg.FindAll([]byte(txt), -1){
		s := string(v[1:len(v)-2])
		rows := strings.Split(s, " = ")
		tmp := regEmphasis.ReplaceAllString(rows[1], "")

		if regInternalLink.FindString(tmp) == ""{
			m[rows[0]] = tmp

		}else if regInternalLink1.FindString(tmp) != "" {
			for _, s := range regInternalLink1.FindAllStringSubmatch(tmp, -1) {
				var res string
				for i := 1; i < len(s); i++ {
					res += s[i]
				}
				res = strings.Replace(res, "[[", "", -1)
				res = strings.Replace(res, "]]", "", -1)
				m[rows[0]] = res
			}
		} else if regInternalLink2.FindString(tmp) != ""{
			for _, s := range regInternalLink2.FindAllStringSubmatch(tmp, -1) {
				var res string
				for i := 1; i < len(s); i++ {
					res += s[i]
				}
				res = strings.Replace(res, "[[", "", -1)
				res = strings.Replace(res, "]]", "", -1)
				m[rows[0]] = res
			}
		}else {
			res := tmp
			res = strings.Replace(res, "[[", "", -1)
			res = strings.Replace(res, "]]", "", -1)
			m[rows[0]] = string(res)
		}
		m[rows[0]] = regExternalLink.ReplaceAllString(m[rows[0]], "")
		m[rows[0]] = regCommentOut.ReplaceAllString(m[rows[0]], "")
	}

	c, err := NewClient("https://en.wikipedia.org/w/api.php")
	if err != nil{
		log.Fatal(err)
	}

	values := url.Values{}
	values.Add("action", "query")
	values.Add("prop", "imageinfo")
	values.Add("format", "json")
	values.Add("iiprop", "url")
	values.Add("titles", "File:" + m["国旗画像"])

	resp, err := http.Get(c.URL.String() + "?" + values.Encode())
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	respjson := ResponceWiki{}
	json.Unmarshal([]byte(b), &respjson)

	fmt.Println(respjson.Query.Pages.Num23473560.Imageinfo[0].URL)
}