package module02

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

// generate a string slice with integer input as size of slice
func generateList(size int) []string {
	list := make([]string, size)
	for i := 0; i < size; i++ {
		list[i] = strconv.Itoa(i)
	}
	return list
}

// TestInterface is a helper for testing functions that sort using
// sort.Interface.
func TestInterface(t *testing.T, sortFn func(sort.Interface)) {
	seed := time.Now().UnixNano()
	t.Log("Seed for random cases:", seed)
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"sorted":           {1, 2, 3, 4},
		"reverse":          {10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"duplicates":       {3, 5, 3, 5, 3, 5},
		"random-len10":     rand.Perm(10),
		"random-len20":     rand.Perm(20),
		"random-len50":     rand.Perm(50),
		"random-len100":    rand.Perm(100),
		"random-len1000":   rand.Perm(1000),
		"random-len10000":  rand.Perm(10000),
		"random-len100000": rand.Perm(100000),
		"sorted-len1000":   intList(1000),
		"sorted-len10000":  intList(10000),
		// "sorted-len100000": intList(100000),
	} {
		t.Run(name, func(t *testing.T) {
			want := make([]int, len(list))
			// for i, val := range list {
			// 	want[i] = val
			// }
			copy(list, want)
			sort.Ints(want)
			sortFn(sort.IntSlice(list))
			errorCount := 0
			if len(list) != len(want) {
				t.Fatalf("got len %d; want %d", len(list), len(want))
			}
			for i := 0; i < len(want); i++ {
				if errorCount >= 5 {
					t.Fatalf("additional errors omitted for brevity")
				}
				if list[i] != want[i] {
					errorCount++
					t.Errorf("list[%d] = %d; want %d", i, list[i], want[i])
				}
			}
		})
	}
}

// BenchmarkInterface is a helper for benchmarking sort functions that sort
// using the sort.Interface.
func BenchmarkInterface(b *testing.B, sortFn func(sort.Interface)) {
	for _, size := range []int{
		100, 200, 400, 800, 1600, 3200, 6400, 12800,
	} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := rand.Perm(size)
				b.StartTimer()
				sortFn(sort.IntSlice(list))
			}
		})
	}
}

// TestPeople is a helper for testing functions that sort people by age then
// last name then first name.
func TestPeopleInterface(t *testing.T, sortFn func([]Person)) {
	for name, list := range map[string][]Person{
		"sorted": {
			{12, "Billy", "Tables"},
			{12, "Bobby", "Tables"},
			{12, "Jordan", "Tables"},
			{12, "Alex", "Zero"},
			{21, "Frank", "Smith"},
			{33, "Bob", "Smilesalot"},
			{45, "Johnny", "Testuser"},
			{65, "Harry", "Hippo"},
			{71, "Thomas", "Train"},
			{53, "Percy", "Engine"},
		},
		"reverse": {
			{71, "Thomas", "Train"},
			{65, "Harry", "Hippo"},
			{53, "Percy", "Engine"},
			{45, "Johnny", "Testuser"},
			{33, "Bob", "Smilesalot"},
			{21, "Frank", "Smith"},
			{12, "Alex", "Zero"},
			{12, "Jordan", "Tables"},
			{12, "Bobby", "Tables"},
			{12, "Billy", "Tables"},
		},
		"duplicates": {
			{53, "Percy", "Engine"},
			{45, "Johnny", "Testuser"},
			{12, "Billy", "Tables"},
			{65, "Harry", "Hippo"},
			{33, "Bob", "Smilesalot"},
			{12, "Bobby", "Tables"},
			{12, "Alex", "Zero"},
			{45, "Johnny", "Testuser"},
			{12, "Billy", "Tables"},
			{21, "Frank", "Smith"},
			{71, "Thomas", "Train"},
			{21, "Frank", "Smith"},
			{12, "Jordan", "Tables"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			want := make([]Person, len(list))
			copy(want, list)
			sort.Sort(people(want))
			sortFn(list)
			errorCount := 0
			if len(list) != len(want) {
				t.Fatalf("got len %d; want %d", len(list), len(want))
			}
			for i := 0; i < len(want); i++ {
				if errorCount >= 5 {
					t.Fatalf("additional errors omitted for brevity")
				}
				if list[i] != want[i] {
					errorCount++
					t.Errorf("list[%d] = %v; want %v", i, list[i], want[i])
				}
			}
		})
	}
}

// TestStringSliceInterface is a helper for testing secret santa function
func TestStringSliceInterface(t *testing.T, secretSantaFn func([]string)) {
	for name, list := range map[string][]string{
		"test with names":           {"Bam Bam", "Barney", "Fred", "Pebbles", "Wilma"},
		"test with reverse names":   {"Wilma", "Pebbles", "Fred", "Barney", "Bam Bam"},
		"test with duplicate names": {"Bam Bam", "Bam Bam", "Bam Bam", "Bam Bam", "Bam Bam"},
		"test with single name":     {"Fred"},
		"test with no names":        {},
	} {
		t.Run(name, func(t *testing.T) {
			secretSantaFn(list)
		})
	}
}

// BenchmarkStringSliceInterface is a helper for benchmarking secret santa function
func BenchmarkStringSliceInterface(b *testing.B, secretSantaFn func([]string)) {
	for _, size := range []int{
		100, 200, 400, 800, 1600, 3200, 6400, 12800,
	} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := generateList(size)
				b.StartTimer()
				secretSantaFn(list)
			}
		})
	}
}
