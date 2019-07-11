package internal

import (
	"testing"
)

func TestContains(t *testing.T) {
	if !Contains([]string{"a", "b"}, "a") {
		t.Errorf("contains should return true when a is in [a,b]")
	}

	if Contains([]string{"a", "b"}, "c") {
		t.Errorf("contains should return false when c is not in [a,b]")
	}
}
