// 101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/

// 给你一个二叉树的根节点 root ，检查它是否轴对称。

// 示例 1：
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true

// 示例 2：
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 时间复杂度O(n) 空间复杂度O(n)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(p, q *TreeNode) bool {
	switch {
	case p == nil || q == nil:
		return p == q
	case p.Val != q.Val:
		return false
	}
	return symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
}
