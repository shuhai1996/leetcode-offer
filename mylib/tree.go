package mylib

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrderTree 前序遍历
func PreOrderTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append([]int{root.Val}, PreOrderTree(root.Left)...), PreOrderTree(root.Right)...)
}

// InOrderTree 中序遍历
func InOrderTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(InOrderTree(root.Left), append([]int{root.Val}, InOrderTree(root.Right)...)...)
}

// PostOrderTree 后序遍历
func PostOrderTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(PostOrderTree(root.Left), PostOrderTree(root.Right)...), root.Val)
}

