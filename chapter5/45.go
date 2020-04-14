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

	newfile, err := os.Create("45.txt")
	if err != nil{
		fmt.Printf("os.Create: %#v", err)
		return
	}
	defer newfile.Close()
	writer := bufio.NewWriter(newfile)

	
	for _, sentence := range sentences{
		for _, chunk := range sentence{
			for _, morph := range chunk.morphs{
				if len(chunk.srcs) == 0 {
					continue
				}
				if morph.pos == "動詞"{
					flg := false
					for _, src := range chunk.srcs{
						for _, morph2 := range sentence[src].morphs{
							if morph2.pos == "助詞"{
								flg = true
							}
						}
					}
					if !flg{
						continue
					}
					_, err := writer.Write([]byte(morph.base))
					if err != nil{
						panic(err)
					}
					_, err = writer.Write([]byte("\t"))
					if err != nil{
						panic(err)
					}
				}else{
					continue
				}
				count := 0
				for _, src := range chunk.srcs{
					for _, morph3 := range sentence[src].morphs{
						if morph3.pos == "助詞"{
							if count > 0 {
									_, err := writer.Write([]byte(" "))
									if err != nil{
										panic(err)
									}
							}
							count++
							_, err := writer.Write([]byte(morph3.base))
							if err != nil{
								panic(err)
							}
						}
					}
				}
				_, err := writer.Write([]byte("\n"))
				if err != nil{
					panic(err)
				}
		    }	
		}
		err := writer.Flush()
		if err != nil{
			panic(err)
		}
	}
}