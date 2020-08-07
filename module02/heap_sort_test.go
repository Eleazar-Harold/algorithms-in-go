package module02

import "testing"

func TestHeapSortInterface(t *testing.T) {
	TestInterface(t, HeapSort)
}

func BenchmarkHeapSortInterface(b *testing.B) {
	BenchmarkInterface(b, HeapSort)
}