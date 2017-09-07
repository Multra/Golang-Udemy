package main

import (
	"fmt"
)

func main() {
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "test1"
	student1[1] = "test2"
	student1[2] = "A"
	student1[3] = "B"
	records = append(records, student1)
	//student 2
	student2 := make([]string, 4)
	student2[0] = "test3"
	student2[1] = "test4"
	student2[2] = "B"
	student2[3] = "B"
	records = append(records, student2)
	fmt.Println(records)

}
