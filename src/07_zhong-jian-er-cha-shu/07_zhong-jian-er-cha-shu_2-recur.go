/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := new(TreeNode)
	if len(preorder) == 0 {
		return nil
	}
	root_val := preorder[0]
	i := 0
	for inorder[i] != root_val {
		i++
	}
	left_tree := buildTree(preorder[1:i+1], inorder[:i])
	right_tree := buildTree(preorder[i+1:], inorder[i+1:])
	root.Val = root_val
	root.Left = left_tree
	root.Right = right_tree
	return root
}