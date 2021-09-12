package bubblesort

import (
	"testing"
)

// TestBubbleSort calls bubblesort.BubbleSort with an array.
// Return sorted array
func TestBubbleSort(t *testing.T) {
	arr := []int{14, 12, 2, 1, 17, 0}
	want := []int{0, 1, 2, 12, 14, 17}
	result := BubbleSort(arr)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result)
		}
	}
}

func TestBubbleSort_RepeatedValue(t *testing.T) {
	arr := []int{2, 0, 2, 1, 1, 0}
	want := []int{0, 0, 1, 1, 2, 2}
	result := BubbleSort(arr)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result)
		}
	}
}

func TestBubbleSort_OneValue(t *testing.T) {
	arr := []int{7}
	want := []int{7}
	result := BubbleSort(arr)
	for i, v := range want {
		if v != result[i] {
			t.Fatalf(`Error %d`, result)
		}
	}
}
