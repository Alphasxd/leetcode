// 144. 二叉树的前序遍历 https://leetcode.cn/problems/binary-tree-preorder-traversal/

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 递归
// func preorderTraversal(root *TreeNode) []int {
//     if root == nil {
//         return nil
//     }
//     var res []int
//     res = append(res, root.Val)
//     res = append(res, preorderTraversal(root.Left)...)
//     res = append(res, preorderTraversal(root.Right)...)
//     return res
// }

// 非递归
func preorderTraversal(root *TreeNode) []int {
    var res []int
    var stack []*TreeNode
    for curr := root; curr != nil || len(stack) > 0; {
        // 沿着左子树一直往下走，直到走到叶子节点
        for curr != nil {
            // 前序遍历，先访问根节点
            res = append(res, curr.Val)
            stack = append(stack, curr)
            curr = curr.Left
        }
        // 栈顶元素出栈，并访问该节点，由于已经访问过左子树，所以直接访问右子树
        curr = stack[len(stack)-1].Right
        stack = stack[:len(stack)-1]
    }
    return res
}
