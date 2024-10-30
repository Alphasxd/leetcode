//124. 二叉树中的最大路径和 https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

//二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。
//该路径 至少包含一个 节点，且不一定经过根节点。路径和 是路径中各节点值的总和。
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
// -1000 <= Node.val <= 1000

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 初始化最大路径和为根节点的值
	maxSum := -1 << 31
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		// 更新最大路径和，即当前节点的值加上左右子树的最大路径和
		maxSum = max(maxSum, root.Val+left+right)
		// 返回的是以当前节点为根节点的最大路径和
		return max(left, right) + root.Val
	}
	dfs(root)
	return maxSum
}
