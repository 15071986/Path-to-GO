package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{4, 2, 3, 1}
	b := []int{1, 2, 3, 4}

	InsertionSort(a)
	if len(a) != len(b) {
		t.Error("Cлайсы имеют разную длину")
	}
	for i, v := range a {
		if v != b[i] {
			t.Errorf("wrong element value on %d index: expected %d, got %d", i, b[i], v)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	result := []int{4, 2, 3, 1}
	for i := 0; i < b.N; i++ {
		InsertionSort(result)
	}
}
