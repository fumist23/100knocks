package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"io"
	"sort"
	"gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
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
type Values []float64

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

	m := sortedKeys(freq)

	p, err := plot.New()
	if err != nil{
		panic(err)
	}

	p.Title.Text = "The order of number"
	p.Y.Label.Text = "count"
	nums := plotter.Values{}
	for _, v := range m[:10]{
		nums = append(nums, float64(freq[v]))
	}
	breath := vg.Points(15)
	bar, err := plotter.NewBarChart(nums, breath)
	if err != nil{
		panic(err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(0)
	p.Add(bar)
	p.Legend.Add("number", bar)
	p.Legend.Top = true
	//p,NominalXは日本語対応してない？？
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "37.png"); err != nil{
		panic(err)
	}

}