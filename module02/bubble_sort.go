package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		swapped := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	for i := 0; i < len(list); i++ {
		swapped := false
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
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
		swapped := false
		for j := 0; j < len(people)-1-i; j++ {
			if less(people[j+1], people[j]) {
				people[j], people[j+1] = people[j+1], people[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSort sorts the given sortable item using Bubble Sort algorithm
// https://en.wikipedia.org/wiki/Bubble_sort
func BubbleSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		swapped := false
		for j := 0; j < list.Len()-1-i; j++ {
			if list.Less(j+1, j) {
				list.Swap(j, j+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
