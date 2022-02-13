package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 再帰を用いた解き方
func preorderTraversal(root *TreeNode) []int {
	var ans []int
	if root != nil {
		// root -> left -> right
		ans = append(ans, root.Val)
		if root.Left != nil {
			ans = append(ans, preorderTraversal(root.Left)...)
		}
		if root.Right != nil {
			ans = append(ans, preorderTraversal(root.Right)...)
		}
	}
	return ans
}

// 配列(slice)を用いた解き方
func preorderTraversalStack(root *TreeNode) []int {
	var ans []int

	if root != nil {
		var stack []*TreeNode
		stack = append(stack, root)

		for len(stack) != 0 {
			lastNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, lastNode.Val)

			if lastNode.Right != nil {
				stack = append(stack, lastNode.Right)
			}

			if lastNode.Left != nil {
				stack = append(stack, lastNode.Left)
			}
		}
	}

	return ans
}
