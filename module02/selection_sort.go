package module02

import "sort"

// SelectionSortInt will sort a list of integers using the selection sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func SelectionSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		lowestIndex := i

		for j := i + 1; j < len(list); j++ {
			if list[j] < list[lowestIndex] {
				lowestIndex = j
			}
		}
		list[i], list[lowestIndex] = list[lowestIndex], list[i]
	}
}

// SelectionSortString is a bubble sort for string slices. Try implementing it for
// practice.
func SelectionSortString(list []string) {
	for i := 0; i < len(list); i++ {
		lowestIndex := i

		for j := i + 1; j < len(list); j++ {
			if list[j] < list[lowestIndex] {
				lowestIndex = j
			}
		}
		list[i], list[lowestIndex] = list[lowestIndex], list[i]
	}
}

// SelectionSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func SelectionSortPerson(people []Person) {
	less := func(a, b Person) bool {
		if a.Age != b.Age {
			return a.Age < b.Age
		}
		if a.LastName != b.LastName {
			return a.LastName < b.LastName
		}
		return a.FirstName < b.FirstName
	}

	for i := 0; i < len(people); i++ {
		lowestIndex := i

		for j := i + 1; j < len(people); j++ {
			if less(people[lowestIndex], people[j]){
				lowestIndex = j
			}
		}
		people[i], people[lowestIndex] = people[lowestIndex], people[i]
	}
}

// SelectionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func SelectionSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		lowestIndex := i

		for j := i + 1; j < list.Len(); j++ {
			if list.Less(j, lowestIndex) {
				lowestIndex = j
			}
		}
		list.Swap(i, lowestIndex)
	}
}
