// 19. 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 示例 1：
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

// 示例 2：
// 输入：head = [1], n = 1
// 输出：[]

// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：O(n)，空间复杂度：O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建一个虚拟头节点，指向首元节点的前一个节点
	dummy := &ListNode{Next: head}

	// 将 head 指针移动到第 n 个节点
	for range n {
		head = head.Next
	}

	// 创建一个前驱节点
	prev := dummy

	// 遍历链表，直到 head 指针指向 nil，此时 prev 指针指向倒数第 n 个节点的前一个节点
	for head != nil {
		head = head.Next
		prev = prev.Next
	}

	//  prev.Next 指向 prev.Next.Next即可删除节点
	prev.Next = prev.Next.Next

	return dummy.Next
}
