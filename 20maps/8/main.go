package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701.txt")
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
	for i := 0; i < 12; i++ {
		fmt.Println(i, len(buckets[i]))
	}
	fmt.Println(buckets[6][3])

}

func HashBucket(word string, buckets int) int {
	hash := 0
	for _, v := range word {
		hash += int(v)
	}
	return hash % buckets
}
