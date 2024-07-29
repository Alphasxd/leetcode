// 83. 删除排序链表中的重复元素 https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

// 给定一个已排序的链表的头 head， 删除所有重复的元素，使得每个元素只出现一次。返回同样按升序排列的结果链表。

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for cur := head; cur != nil && cur.Next != nil; {
		// 如果当前节点和下一个节点的值相等，那么就删除下一个节点
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
