package depthfirstsearch

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(root *TreeNode, searchval int) TreeNode {
	var stack StackNode
	return search(root, &stack, searchval)
}

func search(node *TreeNode, stack *StackNode, searchval int) TreeNode {
	if node.Val == searchval {
		return *node
	}
	if node.Right != nil {
		if !stack.IsExist(node.Right) {
			stack.Push(node.Right)
		}
	}
	if node.Left != nil {
		if !stack.IsExist(node.Left) {
			return search(node.Left, stack, searchval)
		}
	}
	val, exist := stack.Pop()
	if exist {
		return search(&val, stack, searchval)
	}
	var res TreeNode
	return res
}

type StackNode []TreeNode

func (s *StackNode) Push(node *TreeNode) {
	*s = append(*s, *node)
}

func (s *StackNode) Pop() (TreeNode, bool) {
	if len(*s) == 0 {
		var ele TreeNode
		return ele, false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *StackNode) IsExist(node *TreeNode) bool {
	for _, v := range *s {
		if v.Val == node.Val {
			return true
		}
	}
	return false
}
