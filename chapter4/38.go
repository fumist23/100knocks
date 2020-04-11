package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"sort"
	"strconv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
)

type sortedMap struct{
	m map[int]int
	s []int
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

func sortedKeys(m map[int]int)[]int{
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]int, len(m))
	i := 0
	for key, _ := range m{
		sm.s[i] = key
		i++
	}

	sort.Sort(sm)
	return sm.s
}

func drawBarChart(val []float64, label []string){
	group := plotter.Values(val)

	p, err := plot.New()
	if err != nil{
		panic(err)
	}
	p.Title.Text = "the number of frequency"
	p.X.Label.Text = "freauency"
	p.Y.Label.Text = "the number of types"

	w := vg.Points(1)

	bars, err := plotter.NewBarChart(group, w)
	if err != nil{
		panic(err)
	}

	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)

	p.Add(bars)
	p.NominalX(label...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "38.png"); err != nil{
		panic(err)
	}
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

	hist := make(map[int]int)

	for _, count := range freq{
		hist[count]++
	}

	res := sortedKeys(hist)
	var val []float64
	var label []string
	for _, v := range res{
		val = append(val, float64(hist[v]))

		if v == 10 || v == 30 || v == 100{
			label = append(label, strconv.Itoa(v))
		}else{
			label = append(label, "")
		}
	}
	drawBarChart(val, label)

}