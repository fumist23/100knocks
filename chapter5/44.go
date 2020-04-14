package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
	"github.com/awalterschulze/gographviz"

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
	
	n := 10
	g := gographviz.NewGraph()
	if err := g.SetName("G"); err != nil{
		panic(err)
	}
	if err := g.SetDir(true); err != nil{
		panic(err)
	}
	if err := g.AddAttr("G", "bgcolor", "\"#343434\""); err != nil{
		panic(err)
	}

	nodeAttrs := make(map[string]string)
	nodeAttrs["colorscheme"] = "rdylgn11"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontsize"] = "16"
	nodeAttrs["fontcolor"] = "6"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "7"
	nodeAttrs["fillcolor"] = "11"
	nodeAttrs["shape"] = "doublecircle"

	for _, chunk := range sentences[n]{ //10文目で試す
		var word string
		for _, morph := range chunk.morphs{
			if morph.pos != "記号"{
				word += morph.surface
			}
		}
		if err := g.AddNode("G", word, nodeAttrs); err != nil{
			panic(err)
		}
	}

	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"

	for _, chunk := range sentences[10]{
		var wordfrom, wordto string
		for _, morph := range chunk.morphs{
			if morph.pos != "記号"{
				wordfrom += morph.surface
			}
		}
		if chunk.dst != -1{
			for _, morph2 := range sentences[n][chunk.dst].morphs{
				if morph2.pos != "記号"{
					wordto += morph2.surface
				}
			}
			if err := g.AddEdge(wordfrom, wordto, true, edgeAttrs); err != nil{
				panic(err)
			}
		}
	}

	s := g.String()
	file2, err := os.Create("44.dot")
	if err != nil{
		fmt.Printf("os.Create: %#v")
		return
	}
	defer file2.Close()
	file2.Write([]byte(s))
	
}