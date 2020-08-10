package main

import "testing"

// TODO run a series of exponentially increasing test uints
func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(10230498230)
	}
}
func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(10230498230)
	}
}
