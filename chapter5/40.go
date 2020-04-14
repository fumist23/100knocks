package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
)

type Morph struct {
	surface, base, pos, pos1 string
}

func main(){
	filename := "neko.txt.cabocha"
	file, err := os.Open(filename)
	if err != nil{
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sentence := make([]Morph, 0)
	sentences := make([][]Morph, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		txt := string(line)
		if string(line[0]) != "*"{
			if txt != "EOS"{
				txt = strings.Replace(txt, "\t", ",", -1)
				elements := strings.Split(txt, ",")
				m := Morph{elements[0], elements[7], elements[1], elements[2]}
				sentence = append(sentence, m)
			}else{
				if len(sentence) > 0{
					sentences = append(sentences, sentence)
					sentence = make([]Morph, 0)
				}
			}
		}
	}

	fmt.Println(sentences[2])
}