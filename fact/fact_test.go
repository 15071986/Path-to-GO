package fact

import (
	"testing"
)

func TestFact(t *testing.T) {
	result := Fact(4)

	if result != 24 {
		t.Errorf("for 4 expected 24, got %d", result)
	}
}
