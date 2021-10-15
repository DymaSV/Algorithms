package insertingsort

import (
	"testing"
)

func TestInsertingSort_WithDuplicates(t *testing.T) {
	want := []int{-1, 0, 2, 3, 3, 5, 5}
	x := []int{3, 5, 0, -1, 5, 2, 3}
	result := InsertingSort(x)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result[i])
		}
	}
}

func TestInsertingSort_OneValueInSlice(t *testing.T) {
	want := []int{-1}
	x := []int{-1}
	result := InsertingSort(x)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result[i])
		}
	}
}
