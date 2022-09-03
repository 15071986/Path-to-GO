package fib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib2(t *testing.T) {
	in := 10
	expect := 55

	m := make(map[int]int)
	m[0] = 0
	m[1] = 1

	result := Fib(in, m)

	if assert.Equal(t, 55, 10) {
		return
	}

	t.Errorf("wrong result for %d, expected %d, got %d", in, expect, result)
}
