package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Defining Person struct
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

// Defining PersonSlice as an array slice of type Person
type PersonSlice []*Person

// Implementing the 3 methods required by sort.Sort Interfaces
func (p PersonSlice) Len() int      { return len(p) }
func (p PersonSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PersonSlice) Less(i, j int) bool {
	// LastName, then FirstName, then ID
	if p[i].LastName < p[j].LastName {
		return true
	} else if p[i].LastName == p[j].LastName {
		if p[i].FirstName < p[j].FirstName {
			return true
		} else if p[i].FirstName == p[j].FirstName {
			return p[i].ID < p[j].ID
		} else {
			return false
		}
	} else {
		return false
	}
}

//  Print so to be able to print an object of Person construct type
func (p Person) String() string {
	return fmt.Sprintf("%s, %s : %d", p.LastName, p.FirstName, p.ID)
}

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
// edit: we are exploring here using a random random number

func NewPerson(first, last string, ID int) *Person {
	// TODO
	p := new(Person)
	p.FirstName = first
	p.LastName = last
	p.ID = ID
	return p
}

func main() {
	var people PersonSlice
	// instantiating new people (person by person)
	rand.Seed(time.Now().UnixNano())
	ID := rand.Perm(1000)
	p1 := NewPerson("Juan", "Perez", ID[0])
	p2 := NewPerson("Juan", "Soto", ID[1])
	p3 := NewPerson("Carlos", "Perez", ID[2])

	// Future Refactoring: people = append(people, NewPerson(a,b)) does not work for me)

	people = append(people, p1)
	people = append(people, p2)
	people = append(people, p3)
	//people_slice := people[:]
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)

}
