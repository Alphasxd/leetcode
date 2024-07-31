//530. 二叉搜索树的最小绝对差 https://leetcode.cn/problems/minimum-absolute-difference-in-bst/

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
// 差值是一个正数，其数值等于两值之差的绝对值。

package leetcode

import "math"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
    min := math.MaxInt64
    var prev *TreeNode
    var f func(*TreeNode)
    // 中序遍历，得到的是递增序列
    f = func(root *TreeNode) {
        if root == nil {
            return
        }
        f(root.Left)
        // 因为是递增序列，所以只需要计算当前节点和前一个节点的差值
        if prev != nil && root.Val-prev.Val < min {
            min = root.Val - prev.Val
        }
        prev = root
        f(root.Right)
    }
    f(root)
    return min
}
