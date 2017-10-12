package main

import (
	"fmt"
	"sort"
)

type people []string

func (a people) Len() int {
	return len(a)
}

func (a people) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a people) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}
