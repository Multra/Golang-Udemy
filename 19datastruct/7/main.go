package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	const input = "this is a test, a long test. test test."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
