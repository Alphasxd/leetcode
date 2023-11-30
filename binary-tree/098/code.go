// 98. 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
// 有效 二叉搜索树定义如下：
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

// 示例 1：
// 输入：root = [2,1,3]
// 输出：true

// 示例 2：
// 输入：root = [5,1,4,null,null,3,6]
// 输出：false

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, min, max int) bool {
		// 如果当前节点为空，则返回true，符合二叉搜索树的定义
		if root == nil {
			return true
		}
		// 如果当前节点的值不在[min, max]的范围内，则返回false
		if root.Val <= min || root.Val >= max {
			return false
		}
		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}
	// 由于题目中给出的二叉树的值域为[-2^31, 2^31-1]，所以这里使用int的最大值和最小值作为初始值
	return dfs(root, -1<<63, 1<<63-1)
}