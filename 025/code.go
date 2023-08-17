// 25. K 个一组翻转链表 https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]

// 示例 2：
// 输入：head = [1,2,3,4,5], k = 3
// 输出：[3,2,1,4,5]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	// 先求出链表长度
	var len int
	for curr := head; curr != nil; curr = curr.Next {
		len++
	}

	// 创建哑节点
	dummy := &ListNode{Next: head}

	// pre, end 分别指向每次要翻转的链表的头尾节点
	pre, end := dummy, dummy

	// 循环翻转链表
	for end.Next != nil {
		// 定位 end
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 如果 end == nil，说明剩余节点不足 k 个，不需要翻转
		if end == nil {
			break
		}

		start := pre.Next // 翻转链表的头节点
		next := end.Next  // 翻转链表的尾节点的下一个节点

		end.Next = nil                // 断开链表
		pre.Next = reverseList(start) // 翻转链表

		start.Next = next // start 变成翻转链表的尾节点，连接下一个要翻转的链表的头节点
		pre = start       // pre 指向下一个要翻转的链表之前的节点
		end = pre         // end 和 pre 指向同一个节点
	}

	return dummy.Next

}

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
