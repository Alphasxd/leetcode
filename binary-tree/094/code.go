// 94. 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal

// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

// 示例 1：
// 输入：root = [1,null,2,3]
// 1
//  \
//   2
//  /
// 3
// 输出：[1,3,2]

// 示例 2：
// 输入：root = []
// 输出：[]

// 示例 3：
// 输入：root = [1]
// 输出：[1]

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 时间复杂度O(n) 空间复杂度O(n)
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	res = append(res, inorderTraversal(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)
// 	return res
// }

// 迭代 时间复杂度O(n) 空间复杂度O(n)
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for curr := root; curr != nil || len(stack) > 0; {
		// 沿着左子树一直往下走，直到走到叶子节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 栈顶元素出栈，并访问该节点
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}
