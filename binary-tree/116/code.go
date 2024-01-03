// 116. 填充每个节点的下一个右侧节点指针

// 给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。返回树的根节点。

package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 从每层的最左节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过Next指针遍历这一层节点，为下一层的节点更新Next指针
		for node := leftmost; node != nil; node = node.Next {
			// node作为父节点，更新左子节点的Next指针为右子节点
			node.Left.Next = node.Right
			// 如果node有Next节点，更新右子节点的Next指针为node.Next的左子节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	return root
}
