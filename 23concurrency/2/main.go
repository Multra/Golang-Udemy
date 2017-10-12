package main

import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)

//func init() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	fmt.Println(runtime.NumCPU())
//}

var w sync.WaitGroup

func foo() {
	for i := 1; i < 35; i++ {
		fmt.Println("foo", i)
		time.Sleep(5 * time.Millisecond)
	}
	w.Done()
}

func bar() {
	for i := 1; i < 35; i++ {
		fmt.Println("bar", i)
		time.Sleep(50 * time.Millisecond)
	}
	w.Done()
}

func main() {
	w.Add(2)
	go foo()
	go bar()
	w.Wait()
}
