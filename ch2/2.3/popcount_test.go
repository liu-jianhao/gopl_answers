package popcount_test

import (
	"math/rand"
	"testing"
	"time"

	"gopl.io/ch2/solution/2.3"
)

func BenchmarkPopCount(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCount(uint64(rand.Int()))
	}
}
func BenchmarkPopCountByLoop(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCountByLoop(uint64(rand.Int()))
	}
}

func BenchmarkPopCountByShift(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCountByShifting(uint64(rand.Int()))
	}
}
func BenchmarkPopCountByClearing(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		_ = popcount.PopCountByClearing(uint64(rand.Int()))
	}
}
