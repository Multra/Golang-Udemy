package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	h := []byte{}
	for _, num := range bs {
		if num == 10 {
			fmt.Printf("%s\n", h)
			h = h[:0]
		} else {
			h = append(h, num)
			fmt.Print(num, " ")
		}
	}
	//str := string(bs)
	//fmt.Println(str)
}
