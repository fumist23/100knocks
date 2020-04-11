package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"sort"
	"log"
	"math"
	"errors"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

type sortedMap struct{
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int{
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int)bool{
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int){
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int)[]string{
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m{
		sm.s[i] = key
		i++
	}

	sort.Sort(sm)
	return sm.s
}


func plotScatter(X, Y []float64)error{
	if len(X) != len(Y){
		return errors.New("X and Y should be same length")
	}

	scatterData := make(plotter.XYs, len(X))
	for i, _ := range X{
		scatterData[i].X = X[i]
		scatterData[i].Y = Y[i]
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Zipf's law"
	p.X.Label.Text = "the rank of frequency"
	p.Y.Label.Text = "frequency"

	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(scatterData)
	if err != nil{
		panic(err)
	}

	s.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	s.GlyphStyle.Radius = vg.Points(1)

	p.Add(s)
	p.Legend.Add("scatter", s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "39.png"); err != nil{
		panic(err)
	}

	return nil
}

func main(){
	filename := "neko.txt.mecab"
	file, err := os.Open(filename)
	if err != nil{
		fmt.Printf("os.Open: %#v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sentence := make([]map[string]string, 0)
	sentences := make([][]map[string]string, 0)
	for {
		b, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if string(b) != "EOS"{
			txt := string(b)
			txt = strings.Replace(txt, "\t", ",", -1)
			elements := strings.Split(txt, ",")
			morpheme := make(map[string]string)
			morpheme["surface"] = elements[0]
			morpheme["base"] = elements[7]
			morpheme["pos"] = elements[1]
			morpheme["pos1"] = elements[2]
			sentence = append(sentence, morpheme)
		}else{
			if len(sentence) > 0{
				sentences = append(sentences, sentence)
				sentence = make([]map[string]string, 0)
			}
		}
	}

	freq := make(map[string]int)
	for _, sentence := range sentences{
		for _, morpheme := range sentence{
			freq[morpheme["base"]]++
		}
	}

	res := sortedKeys(freq)
	var X, Y []float64

	for i, v := range res{
		X = append(X, math.Log10(float64(i+1)))
		Y = append(Y, math.Log10(float64(freq[v])))
	}

	err = plotScatter(X, Y)
	if err != nil{
		log.Fatal(err)
	}

}