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

func TestStringHas(t *testing.T) {
	val, ok := StringHas("halo more", []string{
		"more",
	})
	if !ok {
		t.Errorf("StringHas should return true when more is in halo more")
	}

	if val != "more" {
		t.Errorf("StringHas should return 'more' when more is in halo more")
	}

	_, ok = StringHas("halo", []string{
		"more",
	})

	if ok {
		t.Errorf("StringHas should not return true when more is in halo more")
	}
}
