package main

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkUseRandPool(b *testing.B) {
	b.ResetTimer()
	uuid.EnableRandPool()
	for i := 0; i < b.N; i++ {
		uuid.NewString()
	}

}

func BenchmarkNotUseRandPool(b *testing.B) {
	b.ResetTimer()
	uuid.DisableRandPool()
	for i := 0; i < b.N; i++ {
		uuid.NewString()
	}
}
