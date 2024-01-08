// 144. 二叉树的前序遍历 https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	res = append(res, root.Val)
// 	res = append(res, preorderTraversal(root.Left)...)
// 	res = append(res, preorderTraversal(root.Right)...)
// 	return res
// }

// 非递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for curr := root; curr != nil || len(stack) > 0; {
		for curr != nil {
			res = append(res, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
