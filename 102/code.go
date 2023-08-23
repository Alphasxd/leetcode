// 102. 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// 示例 1:
// 输入：root = [3,9,20,null,null,15,7]
//            3
//          /  \
//         9   20
//            /  \
//           15   7
// 输出：[[3],[9,20],[15,7]]

// 示例 2:
// 输入：root = [1]
// 输出：[[1]]

// 示例 3：
// 输入：root = []
// 输出：[]

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS 队列实现 时间复杂度O(n) 空间复杂度O(n)
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	// 用切片模拟队列
	queue := []*TreeNode{root}
	// 遍历队列
	for len(queue) > 0 {
		var level []int
		// 遍历当前层的节点
		for range queue  {
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
		// 将当前层的值存入res
		res = append(res, level)
	}
	return res
}
