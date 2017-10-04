package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.gutenberg.org/files/2600/2600-0.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make([]string, 200)
	for scanner.Scan() {
		word := scanner.Text()
		buckets = append(buckets, word)
	}
	//	var inp string
	for i := 0; i < 2; i++ {
		fmt.Println("Enter a word to find what bucket it would go in and how many times it is in the text")
		//		fmt.Scan(&inp)
		fmt.Println("The word", "a", "is in the text", FindBucket(string(i), buckets), "times")
	}
}

func FindBucket(word string, bucket []string) int {
	c := 0
	for i := 0; i < len(bucket); i++ {
		if bucket[i] == word {
			c++
		}
	}
	return c
}
