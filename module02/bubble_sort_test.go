package module02

import "testing"

func TestBubbleSortInt(t *testing.T) {
	TestInt(t, BubbleSortInt)
}

func BenchmarkBubbleSortInt(b *testing.B) {
	BenchmarkInt(b, BubbleSortInt)
}

func TestBubbleSortString(t *testing.T) {
	TestString(t, BubbleSortString)
}

func TestBubbleSortInterface(t *testing.T) {
	TestInterface(t, BubbleSort)
}

func BenchmarkBubbleSortInterface(b *testing.B) {
	BenchmarkInterface(b, BubbleSort)
}
