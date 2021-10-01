package depthfirstsearch

import (
	"testing"
)

// TestDFS calls depthfirstsearch.DFS with a tree.
//Try to find last value
//
//              1					0 1 2 3 4 5 6 7 8 9
//          2        5				1 0 1 1 0 1 0 0 0 0
//       3    4    8   9			2 0 0 1 1 0 0 0 0 0
//     6   7						3 0 1 0 0 0 1 1 0 0
//		  6  						4 0 1 0 0 0 0 0 0 0
//									5 0 0 0 0 0 0 0 1 1
//									6 0 0 1 0 0 0 1 0 0
//									7 0 0 1 0 0 1 0 0 0
//									8 0 0 0 0 1 0 0 0 0
//									9 0 0 0 0 1 0 0 0 0
func TestDFS(t *testing.T) {
	right4 := TreeNode{Val: 9, Left: nil, Right: nil}
	right2 := TreeNode{Val: 4, Left: nil, Right: nil}
	left4 := TreeNode{Val: 8, Left: nil, Right: nil}
	right1 := TreeNode{Val: 5, Left: &left4, Right: &right4}
	left3 := TreeNode{Val: 6, Left: nil, Right: nil}
	right3 := TreeNode{Val: 7, Left: &left3, Right: nil}
	left2 := TreeNode{Val: 3, Left: &left3, Right: &right3}
	left1 := TreeNode{Val: 2, Left: &left2, Right: &right2}
	tree := TreeNode{Val: 1, Left: &left1, Right: &right1}
	want := right2
	result := DFS(&tree, 4)
	if want.Val != result.Val {
		t.Fatalf(`Error %d`, result.Val)
	}
}
