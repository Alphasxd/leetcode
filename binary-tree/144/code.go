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
// 	var res []int
// 	var preorder func(*TreeNode)
// 	preorder = func(node *TreeNode)  {
// 		if node == nil {
// 			return
// 		}
// 		res = append(res, node.Val)
// 		preorder(node.Left)
// 		preorder(node.Right)
// 	}
// 	preorder(root)
// 	return res
// }

// 非递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 访问根节点，并将其入栈
			res = append(res, root.Val)
			stack = append(stack, root)
			// 访问左子树
			root = root.Left
		}
		if len(stack) > 0 {
			// 弹出栈顶元素，并将其右子树置为当前节点
			root = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
