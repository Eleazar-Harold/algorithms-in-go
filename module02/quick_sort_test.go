package module02

import (
	"testing"
)

func TestQuickSortInterface(t *testing.T) {
	TestInterface(t, QuickSort)
}

func BenchmarkQuickSortInterface(b *testing.B) {
	BenchmarkInterface(b, QuickSort)
}
