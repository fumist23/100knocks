package main 

import (
	"fmt"
	"strings"
	"errors"
)

func main() {
	sentence := "I am an NLPer"
	a,err := sentence_ngram(sentence, 2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(a)

	sentence2 := "I am an NLPer"
	b, err := word_ngram(sentence2, 2)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(b)
	
}

func sentence_ngram(sentence string, n int)([]string, error){
	s := strings.Fields(sentence)
	var ngrams []string
	if len(s) < n {
		err := errors.New("Error: input string's length is less than n value")
		return nil, err
	}
	for i := 0; i < (len(s)-n+1); i++ {
		ngrams = append(ngrams, strings.Join(s[i:i+n], "")) 
	}
	return ngrams, nil
}

func word_ngram(sentence string, n int)([]string, error){
	sentence2 := strings.Split(sentence,"")
	sentence3 := strings.Join(sentence2, "")
	s := strings.Split(sentence3, "")
	var ngram2 []string
	if len(s) < 2 {
		err := errors.New("Error: input string's length is less than n value")
		return nil, err
	}
	for i := 0; i < len(s)-n+1; i++ {
		ngram2 = append(ngram2, strings.Join(s[i:i+n], ""))
	}
	return ngram2, nil

}


