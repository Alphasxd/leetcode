// 114. 二叉树展开为链表 https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

// 示例 1：
// 输入：root = [1,2,5,3,4,null,6]
//           1
//         /   \
//        2     5
//       / \     \
//      3   4     6
// 输出：[1,null,2,null,3,null,4,null,5,null,6]

package leetcode

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func flatten(root *TreeNode) {
    curr := root
    // 通过右指针遍历
    for curr != nil {
        // 提前记录当前节点的右子树
        right := curr.Right
        // 将当前节点的左子树插入到右子树的地方
        curr.Left, curr.Right = nil, curr.Left
        // 将原来的右子树接到当前右子树的最右边节点
        prev := curr
        for prev.Right != nil {
            prev = prev.Right
        }
        // 将原来的右子树接到当前右子树的最右边节点
        prev.Right = right
        // 考虑下一个节点
        curr = curr.Right
    }
}
