// 138. 复制带随机指针的链表 https://leetcode.cn/problems/copy-list-with-random-pointer/

// 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
// 构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
// 例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
// 返回复制链表的头节点。
// 用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
// 你的代码 只 接受原链表的头节点 head 作为传入参数。

package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 利用 map 存储原链表的节点和新链表的节点，然后再遍历一次，将新链表的 next 和 random 指针指向正确的节点
// 时间复杂度：O(n)，空间复杂度：O(n)
func copyRandomList(head *Node) *Node {
	// 链表为空，直接返回
	if head == nil {
		return head
	}
	// 用map存储原链表的节点和新链表的节点
	m := make(map[*Node]*Node)

	curr := head
	// 遍历原链表，将原链表的节点和新链表的节点存储到map中，此时新链表的 next 和 random 指针都为空
	for curr != nil {
		node := &Node{Val: curr.Val}
		m[curr] = node
		curr = curr.Next
	}

	curr = head
	// 再次遍历原链表，将新链表的 next 和 random 指针指向正确的节点
	for curr != nil {
		node := m[curr]
		// 将新链表的 next 和 random 指针指向正确的节点
		node.Next = m[curr.Next]
		node.Random = m[curr.Random]
		curr = curr.Next
	}
	// 最后直接返回map中的头节点，对应的值就是新链表的头节点
	return m[head]
}
