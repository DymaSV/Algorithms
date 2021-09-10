package levelorder

import (
	"testing"
)

// TestGetLevelOrder calls levelorder.GetLevelOrder with a tree.
//Try to find last value
//
//              1
//          2        5
//       3    4    8   9
//     6   7
func TestGetLevelOrder(t *testing.T) {
	right4 := TreeNode{Val: 9, Left: nil, Right: nil}
	right3 := TreeNode{Val: 7, Left: nil, Right: nil}
	right2 := TreeNode{Val: 4, Left: nil, Right: nil}
	left4 := TreeNode{Val: 8, Left: nil, Right: nil}
	right1 := TreeNode{Val: 5, Left: &left4, Right: &right4}
	left3 := TreeNode{Val: 6, Left: nil, Right: nil}
	left2 := TreeNode{Val: 3, Left: &left3, Right: &right3}
	left1 := TreeNode{Val: 2, Left: &left2, Right: &right2}
	tree := TreeNode{Val: 1, Left: &left1, Right: &right1}
	want := [][]int{{1}, {2, 5}, {3, 4, 8, 9}, {6, 7}}
	result := LevelOrder(&tree)
	for i, v := range want {
		if !equalArray(v, result[i]) {
			t.Fatalf(`Error %d`, result)
		}
	}
}

func TestGetLevelOrder_TreeIsNil(t *testing.T) {
	want := [][]int{}
	result := LevelOrder(nil)
	for i, v := range want {
		if !equalArray(v, result[i]) {
			t.Fatalf(`Error %d`, result)
		}
	}
}

func TestGetLevelOrder_OnlyRootVal(t *testing.T) {
	tree := TreeNode{Val: 1, Left: nil, Right: nil}
	want := [][]int{{1}}
	result := LevelOrder(&tree)
	for i, v := range want {
		if !equalArray(v, result[i]) {
			t.Fatalf(`Error %d`, result)
		}
	}
}

func equalArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
