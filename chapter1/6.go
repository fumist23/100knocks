package main 

import (
	"fmt"
	"strings"
)
import set "github.com/deckarep/golang-set"

func main() {
	a := "paraparaparadise"
	b := "paragraph"

	//a,bのbi-gramをそれぞれX,Yとおく
	X := ngram(a, 2)
	Y := ngram(b,2)
	fmt.Println(X)
	fmt.Println(Y)
	
	x := set.NewSet()
	y := set.NewSet()
	for _, v := range X{
		x.Add(v)
	}
	for _, v := range Y{
		y.Add(v)
	}

	//和集合
	union := x.Union(y)
	fmt.Println(union)
	//積集合
	intersection := x.Intersect(y)
	fmt.Println(intersection)
	//差集合
	differenceA := x.Difference(y)
	differenceB := y.Difference(x)
	fmt.Println(differenceA)
	fmt.Println(differenceB)

	//'se'というbi-gramがXおよびYに含まれるかどうか
	fmt.Println(x.Contains("se"))
	fmt.Println(y.Contains("se"))

}

func ngram(s string, n int) []string {
	x := strings.Split(s, "")
	var ngrams []string

	for i := 0; i< len(x)+1-n; i++ {
		ngrams = append(ngrams, strings.Join(x[i:i+n], ""))
	}
	return ngrams
}