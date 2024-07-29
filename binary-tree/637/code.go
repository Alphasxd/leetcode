// 637. 二叉树的层平均值 https://leetcode.cn/problems/average-of-levels-in-binary-tree/

// 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	// 用切片模拟队列
	queue := []*TreeNode{root}
	// 遍历队列
	for len(queue) > 0 {
		// 用于存储当前层的节点的值的和
		var sum float64
		// 当前层的节点个数
		size := len(queue)
		// 遍历当前层的节点
		for range queue {
			// 队首出队
			node := queue[0]
			queue = queue[1:]
			// 将出队元素的值累加
			sum += float64(node.Val)
			// 将出队元素的左右子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将当前层的平均值存入res
		res = append(res, sum/float64(size))
	}
	return res
}
