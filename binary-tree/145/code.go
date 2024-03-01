// 145. 二叉树的后序遍历 https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// 给定一个二叉树，返回它的 后序 遍历。

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// func postorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	res = append(res, postorderTraversal(root.Left)...)
// 	res = append(res, postorderTraversal(root.Right)...)
// 	res = append(res, root.Val)
// 	return res
// }

// 非递归
func postorderTraversal(root *TreeNode) []int {
	var res []int

	var stack []*TreeNode
	for curr := root; curr != nil || len(stack) > 0; curr = curr.Right {
		if curr == nil {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		// 将结果插入到头部，反过来就是前序遍历
		res = append([]int{curr.Val}, res...)
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	return res
}