// 148. 排序链表 https://leetcode.cn/problems/sort-list/

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

// 示例 1：
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]

// 示例 2：
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]

// 示例 3：
// 输入：head = []
// 输出：[]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序 时间复杂度O(nlogn) 空间复杂度O(logn)
func sortList(head *ListNode) *ListNode {
	// 链表为空或者只有一个节点时，直接返回
	if head == nil || head.Next == nil {
		return head
	}

	middle := findMiddle(head)

	// 右侧链表
	r := sortList(middle.Next)
	// 将链表断开
	middle.Next = nil

	// 左侧链表
	l := sortList(head)

	// 递归排序左右两侧链表
	return mergeTwoLists(l, r)

}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

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
