package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root != nil {
		// left -> right -> root
		if root.Left != nil {
			ans = append(ans, postorderTraversal(root.Left)...)
		}
		if root.Right != nil {
			ans = append(ans, postorderTraversal(root.Right)...)
		}
		ans = append(ans, root.Val)
	}
	return ans
}
