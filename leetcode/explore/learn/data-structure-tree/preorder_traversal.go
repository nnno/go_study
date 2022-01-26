package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
