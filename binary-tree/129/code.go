// 129. 求根到叶子节点数字之和 https://leetcode.cn/problems/sum-root-to-leaf-numbers/

// 给你一个二叉树的根节点 root， 树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：
// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123。
// 计算从根节点到叶节点生成的 所有数字之和 。

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, prevNum int) int {
	if root == nil {
		return 0
	}
	sum := prevNum*10 + root.Val
	// 如果当前节点是叶子节点，直接返回sum
	if root.Left == nil && root.Right == nil {
		return sum
	}
	// 如果当前节点不是叶子节点，返回左右子树的和，传入的参数是sum作为下次递归的prevNum
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}
