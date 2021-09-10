package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Level order traversal of a binary tree
func LevelOrder(tree *TreeNode) [][]int {
	if tree == nil {
		return [][]int{}
	}
	var result [][]int = [][]int{{tree.Val}}
	if tree.Left == nil && tree.Right == nil {
		return result
	}
	return getLevelOrder(tree, &result, 1)
}

func getLevelOrder(tree *TreeNode, arr *[][]int, level int) [][]int {
	result := *arr
	oneLevelVal := getValueFromLevel(tree)
	if oneLevelVal != nil {
		if len(result) <= level {
			result = append(result, oneLevelVal)
		} else {
			result[level] = append(result[level], oneLevelVal...)
		}
	}
	level = level + 1
	if tree.Left != nil {
		result = getLevelOrder(tree.Left, &result, level)
	}
	if tree.Right != nil {
		result = getLevelOrder(tree.Right, &result, level)
	}
	return result
}

func getValueFromLevel(tree *TreeNode) []int {
	var oneLevelVal []int
	if tree.Left != nil {
		oneLevelVal = append(oneLevelVal, tree.Left.Val)
	}
	if tree.Right != nil {
		oneLevelVal = append(oneLevelVal, tree.Right.Val)
	}
	return oneLevelVal
}
