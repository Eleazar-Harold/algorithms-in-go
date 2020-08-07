package module02

import "testing"

func TestCocktailSortInterface(t *testing.T) {
	TestInterface(t, CocktailSort)
}

func BenchmarkCocktailSortInterface(b *testing.B) {
	BenchmarkInterface(b, CocktailSort)
}
