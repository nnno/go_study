package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	if root != nil {
		// left -> root -> right
		if root.Left != nil {
			ans = append(ans, inorderTraversal(root.Left)...)
		}
		ans = append(ans, root.Val)
		if root.Right != nil {
			ans = append(ans, inorderTraversal(root.Right)...)
		}
	}
	return ans
}
