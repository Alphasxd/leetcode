// 226. 翻转二叉树 https://leetcode.cn/problems/invert-binary-tree/

// 给你一颗二叉树的根节点 root，翻转这颗二叉树，并返回其根节点。

// 示例 1：
// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]

// 示例 2：
// 输入：root = [2,1,3]
// 输出：[2,3,1]

// 示例 3：
// 输入：root = []
// 输出：[]

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归，时间复杂度O(n)，空间复杂度O(n)
func invertTree(root *TreeNode) *TreeNode {
	
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
