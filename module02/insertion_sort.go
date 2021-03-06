package module02

import "sort"

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	for i := 0; i <= len(list)-1; i++ {
		currentValue := list[i]
		position := i

		for position > 0 && list[position-1] > currentValue {
			list[position], position = list[position-1], position-1
		}

		list[position] = currentValue
	}
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
	for i := 0; i <= len(list)-1; i++ {
		currentValue := list[i]
		position := i

		for position > 0 && list[position-1] > currentValue {
			list[position], position = list[position-1], position-1
		}

		list[position] = currentValue
	}
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
	less := func(a, b Person) bool {
		if a.Age != b.Age {
			return a.Age < b.Age
		}
		if a.LastName != b.LastName {
			return a.LastName < b.LastName
		}
		return a.FirstName < b.FirstName
	}

	for i := 0; i <= len(people)-1; i++ {
		currentValue := people[i]
		position := i

		for position > 0 && less(currentValue, people[position-1]) {
			people[position], position = people[position-1], position-1
		}

		people[position] = currentValue
	}
}

// InsertionSort sorts the given sortable item using Insertion Sort algorithm
// https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(list sort.Interface) {
	for i, j := 1, 1; i < list.Len(); i, j = i+1, i+1 {
		for (j > 0) && list.Less(j, j-1) {
			list.Swap(j, j-1)
			j--
		}
	}
}
