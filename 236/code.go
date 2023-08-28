// 236.二叉树的最近公共祖先 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

// 给定一个二叉树，找到该树中两个指定节点的最近公共祖先。

// 示例 1:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
//             3
//          /     \
//        5         1
//      /   \     /   \
//     6     2   0     8
//          / \
//         7   4

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 若 root 为空，或者 root 等于 p 或者 q，则返回 root
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 递归左子树
	right := lowestCommonAncestor(root.Right, p, q) // 递归右子树

	if left != nil && right != nil { // 若左子树和右子树都不为空，则返回 root
		return root
	}

	if left != nil {
		return left // 若左子树不为空，则返回左子树
	}
	return right // 若右子树不为空，则返回右子树
}