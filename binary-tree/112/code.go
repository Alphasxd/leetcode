// 112. 路径总和 https://leetcode.cn/problems/path-sum/

// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
    // 空树，直接返回false
    if root == nil {
        return false
    }
    // 只存在根节点，判断根节点的值是否等于sum
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    // 递归判断左右子树
    sum -= root.Val
    return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
