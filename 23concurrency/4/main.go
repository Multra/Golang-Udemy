package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func inc(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, counter)
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go inc("foo")
	go inc("bar")
	wg.Wait()
	fmt.Println("Counter", counter)
}
