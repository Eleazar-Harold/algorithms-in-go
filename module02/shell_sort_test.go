package module02

import "testing"

func TestShellSortInterface(t *testing.T) {
	TestInterface(t, ShellSort)
}

func BenchmarkShellSortInterface(b *testing.B) {
	BenchmarkInterface(b, ShellSort)
}
