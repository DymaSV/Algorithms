package binsearchloop

import (
	"testing"
)

// TestBinSearch calls binsearchloop.BinSearchLoop with an array.
//Try to find last value
func TestBinSearchLoop_LastValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := 5
	result := BinSearchLoop(&arr, 13)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Try to find first value
func TestBinSearchLoop_FirstValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := 0
	result := BinSearchLoop(&arr, 2)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Try to find middle value
func TestBinSearchLoop_MiddleValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := 2
	result := BinSearchLoop(&arr, 5)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Did not find value
func TestBinSearchLoop_NotFindValue(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := -1
	result := BinSearchLoop(&arr, 15)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}

//Did not find value in the middle of array
func TestBinSearchLoop_NotFindValueInMiddle(t *testing.T) {
	arr := []int{2, 3, 5, 7, 11, 13}
	want := -1
	result := BinSearchLoop(&arr, 1)
	if want != result {
		t.Fatalf(`Error %d`, result)
	}
}
