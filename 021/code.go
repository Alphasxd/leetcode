// 21. 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists/

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 示例 1：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]

// 示例 2：
// 输入：l1 = [], l2 = []
// 输出：[]

// 示例 3：
// 输入：l1 = [], l2 = [0]
// 输出：[0]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 时间复杂度：O(n+m)，空间复杂度：O(1)
func mergeTwoLists(l1, l2 *ListNode) *ListNode {

	// 创建一个虚拟头节点，在合并后的链表前面，用于返回合并后的链表
	dummy := new(ListNode)
	// 创建一个游标
	cur := dummy
	// 循环比较两个链表的值，将较小的值放入新链表中
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	// 有一个链表为空时，将另一个链表剩余的值放入新链表中
	switch {
	case l1 != nil:
		cur.Next = l1
	case l2 != nil:
		cur.Next = l2
	}

	// 返回辅助节点的下一个节点，即合并后的链表
	return dummy.Next
}
