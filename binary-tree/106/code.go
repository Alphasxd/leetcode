// 106. 从中序与后序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

// 给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    length := len(postorder)
    if length == 0 {
        return nil
    }
    i := func(order []int, v int) int {
        var index int
        for order[index] != v {
            index++
        }
        return index
    }(inorder, postorder[length-1])

    // 中序遍历中的根节点索引为 i
    // 后序遍历中的根节点索引为 length-1
    // 在递归调用的时候注意去除根节点
    return &TreeNode{
        postorder[length-1],
        buildTree(inorder[:i], postorder[:i]),
        buildTree(inorder[i+1:], postorder[i:length-1]),
    }
}
