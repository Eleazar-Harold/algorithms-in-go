package module02

import "sort"

// Person is used in a few sorting practice problems.
type Person struct {
	Age       int
	FirstName string
	LastName  string
}

// Try not to read this code if you want to try to implement the sort.Interface
// on your own.
type people []Person

func (p people) Len() int { return len(p) }
func (p people) Less(i, j int) bool {
	a, b := p[i], p[j]
	if a.Age != b.Age {
		return a.Age < b.Age
	}
	if a.LastName != b.LastName {
		return a.LastName < b.LastName
	}
	return a.FirstName < b.FirstName
}
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// func SortPeople() aims to sort people by age and first name
// in order to avoid collisions. Refer to internal people functions
// to help achieve this
func SortPeople(p []Person) {
	sort.Sort(people(p))
}
