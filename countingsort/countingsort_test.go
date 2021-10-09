package countingsort

import (
	"testing"
)

// TestDFS calls depthfirstsearch.DFS with a tree.
//Try to find last value
func TestCountSort(t *testing.T) {
	// want := 7
	x := []int{5, 3, 5, 4, 1}
	CountSort(x)
	// if want != result {
	// 	t.Fatalf(`Error %d`, result.Val)
	// }
}
