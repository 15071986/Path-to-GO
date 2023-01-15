package main

import (
	"testing"
)

func BenchmarkRLock(b *testing.B) {
	for i := 0; i < 1e3; i++ {
		rw.RLock()
	}
}
func BenchmarkLock(b *testing.B) {
	for i := 0; i < 1e3; i++ {
		rw.Lock()

	}
}
