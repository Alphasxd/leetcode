// 117. 填充每个节点的下一个右侧节点指针 II https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii

// 给定一个二叉树，填充它的每个 next 指针，让这个指针指向其下一个右侧节点。

package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层序遍历，使用辅助队列
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		// 当前层的节点数量
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// 除去最后一个节点，其余节点的 Next 指向下一个节点，也就是队列的第一个节点
			if i < size-1 {
				node.Next = queue[0]
			}

			// 将当前节点的左右子节点入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
