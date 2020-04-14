package main 

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"
)

type Morph struct{
	surface, base, pos, pos1 string
}

type Chunk struct{
	morphs []Morph
	dst int
	srcs []int
}

func gotoNextChunk(sentence []Chunk, current int){
	chunk := sentence[current]
	var word string
	for _, morph := range chunk.morphs{
		if morph.pos != "記号"{
			word += morph.surface
		}
	}

	if chunk.dst != -1{
		fmt.Print(word, "->")
		gotoNextChunk(sentence, chunk.dst)
	}else{
		fmt.Println(word)
	}
}

func main(){
	file, err := os.Open("neko.txt.cabocha")
	if err != nil{
		fmt.Printf("os.Open: %#v", err)
		return
	}

	reader := bufio.NewReader(file)

	sentences := make([][]Chunk, 0)
	sentence := make([]Chunk, 0)
	i := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
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

		for _, sentence := range sentences{
			for _, chunk := range sentence{
				var word string
				var flg bool

				for _, morph := range chunk.morphs{
					if morph.pos != "記号"{
						word += morph.surface
					}
					if morph.pos == "名詞"{
						flg = true
					}
				}

				if flg {
					if chunk.dst != -1{
						fmt.Print(word, "->")
						gotoNextChunk(sentence, chunk.dst)
					}else{
						fmt.Println(word)
					}
				}
			}
		}
	}
}