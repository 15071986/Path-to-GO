package fib

import (
	"testing"
)

func TestMain(t *testing.T) {
	in := 10
	expect := 55

	m := make(map[int]int)
	m[0] = 0
	m[1] = 1

	result := Fib(in, m)

	if result != expect {
		t.Errorf("wrong result for %d, expected %d, got %d", in, expect, result)
	}
}
