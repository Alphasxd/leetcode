// 543. 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/

// 给你一棵二叉树的根节点，返回该树的 直径 。
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的 长度 由它们之间边数表示。

// 示例 1：
// 输入：root = [1,2,3,4,5]
//      1
//     / \
//    2   3
//   / \
//  4   5
// 输出：3

// 示例 2：
// 输入：root = [1,2]
// 输出：1

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 递归 时间复杂度O(n) 空间复杂度O(n)
func diameterOfBinaryTree(root *TreeNode) int {
    var res int
    var diameter func(*TreeNode) int

    diameter = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        // 递归计算左右子树的深度
        left := diameter(root.Left)
        right := diameter(root.Right)
        // 计算当前节点的最大直径
        if left+right > res {
            res = left + right
        }

        // 返回当前节点的深度
        depth := left
        if right > depth {
            depth = right
        }
        return depth + 1
    }
    // 计算二叉树的直径
    diameter(root)
    return res
}
