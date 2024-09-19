// 143. 重排链表 https://leetcode-cn.com/problems/reorder-list/

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

// L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为：

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 示例 1:
// 输入：head = [1,2,3,4]
// 输出：[1,4,2,3]

// 示例 2:
// 输入：head = [1,2,3,4,5]
// 输出：[1,5,2,4,3]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：O(n)，空间复杂度：O(1)
func reorderList(head *ListNode) {
	// 找到链表中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半部分链表
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}
	// 合并两个链表
	for p, q := head, prev; p != q; p, q = q, p {
		next := p.Next
		p.Next = q
		p = next
	}
}
