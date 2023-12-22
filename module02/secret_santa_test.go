package module02

import (
	"testing"
)

// TestSecretSanta tests the SecretSanta function
func TestSecretSantaInterface(t *testing.T) {
	TestStringSliceInterface(t, SecretSanta)
}

// BenchmarkSecretSanta with multiple scenarios
func BenchmarkSecretSanta(b *testing.B) {
	BenchmarkStringSliceInterface(b, SecretSanta)
}
