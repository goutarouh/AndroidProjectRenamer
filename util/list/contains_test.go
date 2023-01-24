package list

import "testing"

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c"}

	if !Contains(s, "a") {
		t.Fatal("Contains_test errror", "a not in ", s)
	}

	if Contains(s, "d") {
		t.Fatal("Contains_test error", "d in ", s)
	}
}
