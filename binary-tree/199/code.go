// 199. 二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

// 示例 1:
// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]

// 示例 2:
// 输入: [1,null,3]
// 输出: [1,3]

// 示例 3:
// 输入: []
// 输出: []

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 使用dfs，深度优先搜索，并记录深度，先遍历右子树，再遍历左子树
func rightSideView(root *TreeNode) []int {
    var res []int
    var dfs func(*TreeNode, int)
    dfs = func(tn *TreeNode, depth int) {
        if tn == nil {
            return
        }
        // 如果深度 = 结果切片的长度，就将遍历到的节点添加到切片中
        // 因为是先遍历右子树，所以右子树的节点会先添加到切片中
        // 因为二叉树的 depth 是从 0 开始的，所以当 depth = len(res) 时，说明是第一个遍历到的节点
        if depth == len(res) {
            res = append(res, tn.Val)
        }
        // 递归右子树，深度+1，递归左子树，深度+1
        dfs(tn.Right, depth+1)
        dfs(tn.Left, depth+1)
    }
    dfs(root, 0)
    return res
}
