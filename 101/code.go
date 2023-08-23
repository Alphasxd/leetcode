// 101. 对称二叉树 https://leetcode-cn.com/problems/symmetric-tree/

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
	var check func(*TreeNode, *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root, root)
}
