
//理解できず、、、、、

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

	newfile, err := os.Create("47.txt")
	if err != nil{
		fmt.Printf("os.Create: %#v", err)
		return
	}
	defer newfile.Close()
	writer := bufio.NewWriter(newfile)

	
	for _, sentence := range sentences{
		for _, chunk := range sentence{
			for _, morph := range chunk.morphs{
				if len(chunk.srcs) == 0{
					continue
				}
				if morph.pos == "動詞"{
					flg := false
					var predicate string
					wo := -1
					count := 0
					for _, src := range chunk.srcs{
						wo = src
						for i, morph2 := range sentence[src].morphs{
							if morph2.pos == "助詞"{
								if morph2.surface == "を"{
									if sentence[src].morphs[i-1].pos == "名詞" && sentence[src].morphs[i-1].pos1 == "サ変接続"{
										flg = true
										predicate = sentence[src].morphs[i-1].base + morph2.base + morph.base
									}
								}
								count++
							}
						}
					}
					if !flg {
						continue
					}else if count == 1{
						continue
					}

					_, err := writer.Write([]byte(predicate))
					if err != nil{
						panic(err)
					}

					_, err = writer.Write([]byte("\t"))
					if err != nil{
						panic(err)
					}

					count = 0
					for _, src := range chunk.srcs{
						last_post := -1
						for i, morph2 := range sentence[src].morphs{
							if morph2.pos == "助詞" && src != wo{
								last_post = i
							}
						}
						if count > 0{
							_, err = writer.Write([]byte(sentence[src].morphs[last_post].surface))
							if err != nil{
								panic(err)
							}
						}
					}

					_, err = writer.Write([]byte("\t"))
					if err != nil{
						panic(err)
					}

					count = 0
					for _, src := range chunk.srcs{
						var word string
						last_post := -1
						for i, morph2 := range sentence[src].morphs{
							if morph2.pos != "記号"{
								word += morph.surface
							}
							if morph2.pos == "動詞" && src != wo{
								last_post = i
							}
							if count > 0{
								_, err := writer.Write([]byte(" "))
								if err != nil{
									panic(err)
								}
							}
							count++
						}
						if last_post != -1{
							_, err := writer.Write([]byte(word))
							if err != nil{
								panic(err)
							}
						}
					}
					_, err = writer.Write([]byte("\n"))
					if err != nil{
						panic(err)
					}
					if flg{
						break
					}
				}
			}
		}
		err := writer.Flush()
		if err != nil{
			panic(err)
		}
	}
}