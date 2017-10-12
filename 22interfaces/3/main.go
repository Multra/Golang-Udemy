package main

import (
	"fmt"
	"sort"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

func (a PersonSlice) Len() int { return len(a) }
func (a PersonSlice) Less(i, j int) bool {
	if a[i].LastName != a[j].LastName {
		return a[i].LastName < a[j].LastName
	} else if a[i].LastName == a[j].LastName {
		return a[i].FirstName < a[j].FirstName
	} else if a[i].FirstName == a[j].FirstName {
		return a[i].ID < a[j].ID
	} else {
		return false
	}
}
func (a PersonSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Person) String() string {
	return fmt.Sprintf("ID: %d, Last: %s, First: %s", a.ID, a.LastName, a.FirstName)
}

func main() {
	f, l := "", ""
	z := PersonSlice{}
	for c := 0; c < 3; c++ {
		fmt.Println("Enter first name then last name")
		fmt.Scan(&f)
		fmt.Scan(&l)
		z = append(z, NewPerson(len(z), f, l))
	}
	sort.Sort(z)
	fmt.Println(z)
}

func NewPerson(ID int, first, last string) *Person {
	p := new(Person)
	p.ID = ID + 1
	p.FirstName, p.LastName = first, last
	return p
}
