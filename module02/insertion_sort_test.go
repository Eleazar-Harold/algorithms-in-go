package module02

import (
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	TestInt(t, InsertionSortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	BenchmarkInt(b, InsertionSortInt)
}

func TestInsertionSortString(t *testing.T) {
	TestString(t, InsertionSortString)
}

func TestInsertionSortInterface(t *testing.T) {
	TestInterface(t, InsertionSort)
}
