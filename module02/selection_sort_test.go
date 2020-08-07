package module02

import "testing"

func TestSelectionSortInt(t *testing.T) {
	TestInt(t, SelectionSortInt)
}

func BenchmarkSelectionSortInt(b *testing.B) {
	BenchmarkInt(b, SelectionSortInt)
}

func TestSelectionSortString(t *testing.T) {
	TestString(t, SelectionSortString)
}

func TestSelectionSortInterface(t *testing.T) {
	TestInterface(t, SelectionSort)
}

func BenchmarkSelectionSortInterface(b *testing.B) {
	BenchmarkInterface(b, SelectionSort)
}
