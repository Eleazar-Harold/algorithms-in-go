package module02

import "sort"

// Sort function will combine
// left and right slices into a newly created slice/list.
func Sort(s sort.Interface) {
	list := make([]int, s.Len())
	var mergeSort func(start, end int)
	mergeSort = func(start, end int) {
		size := end - start
		middle := start + (size / 2)
		switch {
		case size < 2:
			return
		}
		mergeSort(start, middle)
		mergeSort(middle, end)
		li, ls := start, middle
		ri, rs := middle, end
		cursor := 0
		for ; li < ls && ri < rs; cursor++ {
			if s.Less(li, ri) {
				list[li-start] = cursor
				li++
			} else {
				list[ri-start] = cursor
				ri++
			}
		}
		for ; li < ls; li++ {
			list[li-start] = cursor
			cursor++
		}
		for i := range list[:cursor] {
			for j := list[i]; j != i; {
				s.Swap(start+i, start+j)
				list[i], list[j], j = list[j], list[i], list[j]
			}
		}
	}
	mergeSort(0, s.Len())
}

// MergeSort sorts the given sortable item using MergeSort Sort algorithm
// https://en.wikipedia.org/wiki/Merge_sort
func MergeSort(s sort.Interface) {
	// If the list is a single element, return it
	if s.Len() < 2 || s == nil {
		return
	}
	Sort(s)
}
