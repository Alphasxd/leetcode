// 222. 完全二叉树的节点个数 https://leetcode-cn.com/problems/count-complete-tree-nodes/

// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var cnt int
	// 根节点为第 0 层，所以树的最高层为 height - 1
	for h := height(root) - 1; root != nil; h-- {
		// 右子树的高度
		rh := height(root.Right)

		// 如果右子树的高度等于 h(height - 1)，说明左子树是满二叉树，左子树的节点个数为 2^h，加上根节点，再加上右子树的节点个数
		if h == rh {
			cnt += 1 << uint(h)
			root = root.Right
			h = rh
		} else { // 如果右子树的高度不等于 h，说明右子树是满二叉树，右子树的节点个数为 2^rh，加上根节点，再加上左子树的节点个数
			cnt += 1 << uint(rh)
			root = root.Left
		}
	}
	return cnt
}

// 计算完全二叉树的高度，一直向左走到底，即为高度
func height(root *TreeNode) int {
	var h int
	for p := root; p != nil; p = p.Left {
		h++
	}
	return h
}