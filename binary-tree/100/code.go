// 100. 相同的树 https://leetcode.cn/problems/same-tree/

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil || q == nil: // 如果 p 和 q 中有一个是 nil，则返回 p == q
		return p == q
	case p.Val != q.Val: // 如果 p 和 q 的值不相等，则返回 false
		return false
	}
	// 递归判断 p 和 q 的左右子树是否相同
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
