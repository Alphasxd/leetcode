// 103. 二叉树的锯齿形层序遍历 https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	// 用切片模拟队列
	queue := []*TreeNode{root}
	// 遍历队列
	zigzag := 0
	for len(queue) > 0 {
		var level []int
		// 遍历当前层的节点
		for range queue {
			// 队首出队
			node := queue[0]
			queue = queue[1:]
			// 将出队元素的值存入level，记录当前层的值
			level = append(level, node.Val)
			// 将出队元素的左右子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if zigzag%2 == 1 {
			for i := 0; i < len(level)/2; i++ {
				level[i], level[len(level)-1-i] = level[len(level)-1-i], level[i]
			}
		}
		zigzag++
		// 将当前层的值存入res
		res = append(res, level)
	}
	return res
}
