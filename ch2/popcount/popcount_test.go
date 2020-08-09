package popcount

import "testing"

func BenchmarkPopCount1(b *testing.B) {
	PopCount1(10230498230)
}
func BenchmarkPopCount2(b *testing.B) {
	PopCount2(10230498230)
}
