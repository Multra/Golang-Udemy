package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func inc(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, counter)
		mutex.Unlock()
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
