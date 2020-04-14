package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
)

type Morph struct{
	surface, base, pos, pos1 string
}

type Chunk struct{
	morphs []Morph
	dst int
	srcs []int
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
	sentence := make([]Chunk, 0)
	sentences := make([][]Chunk, 0)
	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		txt := string(line)
		if txt != "EOS"{
			if string(line[0]) == "*"{
				elements := strings.Split(txt, "D")
				elements = strings.Split(elements[0], " ")
				dst, _ := strconv.Atoi(elements[2])
				c := Chunk{dst: dst}
				sentence = append(sentence, c)
				i++
			}else{
				txt = strings.Replace(txt, "\t", ",", -1)
				elements := strings.Split(txt, ",")
				m := Morph{elements[0], elements[7], elements[1], elements[2]}
				sentence[i-1].morphs = append(sentence[i-1].morphs, m)
			}
		}else{
			if len(sentence) > 0{
				for i, chunk := range sentence{
					if chunk.dst != -1{
						sentence[chunk.dst].srcs = append(sentence[chunk.dst].srcs, i)
					}
				}
				sentences = append(sentences, sentence)
				sentence = make([]Chunk, 0)
			}
			i = 0
		}
		
	}
	for _, sentence := range sentences{
		for _, chunk := range sentence{
			if chunk.dst != -1{
				flg := true
				for i, morph := range chunk.morphs{
					if morph.pos == "名詞"{
						break
					}else if i == len(chunk.morphs)-1{
						flg = false
					}
				}

				for i, morph2 := range sentence[chunk.dst].morphs{
					if morph2.pos == "動詞"{
						break
					}else if i == len(sentence[chunk.dst].morphs)-1 {
						flg = false
					}
				}

				if flg == false{
					continue
				}

				for _, morph := range chunk.morphs{
					if morph.pos != "記号"{
						fmt.Print(morph.surface)
					}
				}
				fmt.Print("\t")
				for _, morph2 := range sentence[chunk.dst].morphs{
					if morph2.pos != "記号"{
						fmt.Print(morph2.surface)
					}
				}
				fmt.Println("")
			}
		}
	}

	
}