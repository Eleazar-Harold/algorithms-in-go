package module02

import "testing"

func TestMergeSortInterface(t *testing.T) {
	TestInterface(t, MergeSort)
}

func BenchmarkMergeSortInterface(b *testing.B) {
	BenchmarkInterface(b, MergeSort)
}
