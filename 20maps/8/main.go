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
	buckets := make([][]string, 12)
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	//	for i := 0; i < 12; i++ {
	//		fmt.Println(i, len(buckets[i]))
	//	}

	//	var inp string
	for i := 0; i < 2; i++ {
		fmt.Println("Enter a word to find what bucket it would go in and how many times it is in the text")
		//		fmt.Scan(&inp)
		wordBucket := HashBucket(string(i), 12)
		fmt.Println("a", "is in bucket", wordBucket)
		fmt.Println("The word", "a", "is in the text", FindBucket(string(i), buckets[wordBucket]), "times")
	}
}

func HashBucket(word string, buckets int) int {
	hash := 0
	for _, v := range word {
		hash += int(v)
	}
	return hash % buckets
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
