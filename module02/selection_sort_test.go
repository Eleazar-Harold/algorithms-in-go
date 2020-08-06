package module02

import "testing"

func TestSelectionSortInt(t *testing.T) {
	TestInt(t, SelectionSortInt)
}

func BenchmarkSelectionSortInt(b *testing.B) {
	BenchmarkInt(b, SelectionSortInt)

	// Below is the code used in the video. I have opted to use a single function
	// in the sorttest package for all benchmarks moving forward to make it easier
	// to benchmark all of the sort functions we cover.
	//
	// for _, size := range []int{
	// 	100, 200, 400, 800, 1600,
	// } {
	// 	b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
	// 		for n := 0; n < b.N; n++ {
	// 			b.StopTimer()
	// 			list := rand.Perm(size)
	// 			b.StartTimer()
	// 			SelectionSortInt(list)
	// 		}
	// 	})
	// }
}

func TestSelectionSortString(t *testing.T) {
	TestString(t, SelectionSortString)
}

func TestSelectionSortInterface(t *testing.T) {
	TestInterface(t, SelectionSort)
}
