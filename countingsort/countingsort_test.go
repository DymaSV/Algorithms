package countingsort

import (
	"testing"
)

// TestCountingSort calls countingsort.CountingSort.
func TestCountingSort_WithDuplicates(t *testing.T) {
	want := []int{0, 1, 2, 3, 3, 5, 5}
	x := []int{3, 5, 0, 1, 5, 2, 3}
	result := CountingSort(x)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result[i])
		}
	}
}

func TestCountingSort_WithOutDuplicates(t *testing.T) {
	want := []int{0, 1, 2, 3, 5, 12, 16}
	x := []int{3, 5, 0, 1, 12, 2, 16}
	result := CountingSort(x)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result[i])
		}
	}
}

func TestCountingSort_OneValue(t *testing.T) {
	want := []int{0}
	x := []int{0}
	result := CountingSort(x)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result[i])
		}
	}
}

func TestCountingSort_EmptyArray(t *testing.T) {
	want := []int{}
	x := []int{}
	result := CountingSort(x)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result[i])
		}
	}
}
