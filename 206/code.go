// 206.反转链表 https://leetcode-cn.com/problems/reverse-linked-list/

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]

// 示例 3：
// 输入：head = []
// 输出：[]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// head 是指向首元结点的指针，leetcode 上的链表都没有头结点
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
	for head != nil {
		next := head.Next // next 指向 curr 的下一个节点
		head.Next = prev  // 反转 curr 的指针
		prev = head       // prev 指向 curr
		head = next       // curr 指向 next
	}
    // prev 指向反转后的链表的头节点
    return prev
}