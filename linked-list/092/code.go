// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

// 示例 1：
// 输入：head = [1,2,3,4,5], left = 2, right = 4
// 输出：[1,4,3,2,5]

// 示例 2：
// 输入：head = [5], left = 1, right = 1
// 输出：[5]

package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 哑节点，因为考虑到 left = 1 的情况，所以需要一个哑节点，统一操作
	dummy := &ListNode{Next: head}
	pre := dummy
	// 将 pre 移动到 left 的前一个节点
	for i := 0; i < left - 1; i++ {
		pre = pre.Next
	}
	// 从 left 开始，头插法插入到 pre 后面
	curr := pre.Next
	for i := left; i < right; i++ {
		// 防止断链
		next := curr.Next
		// 头插法
		curr.Next = next.Next
		// next 指向 pre 的下一个节点
		next.Next = pre.Next
		// pre 指向 next
		pre.Next = next
	}
	return dummy.Next
}