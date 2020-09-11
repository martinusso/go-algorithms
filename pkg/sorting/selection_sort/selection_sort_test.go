package selection_sort

import (
	"testing"
)

var shuffled = []int{1, 8, 6, 2, 10, 5, 4, 7, 9, 3}

func TestSelectionSortString(t *testing.T) {
	got := Sort(shuffled)
	if len(shuffled) != len(got) {
		t.Errorf("Expected '%d', got '%d'", len(shuffled), len(got))
	}
	for i := 0; i < len(got); i++ {
		if i+1 != got[i] {
			t.Errorf("Incorrect order for '%d' ('%d')", got[i], i)
		}
	}
}
