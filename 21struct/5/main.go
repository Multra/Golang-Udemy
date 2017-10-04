package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	json.NewDecoder(strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}}`)).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
