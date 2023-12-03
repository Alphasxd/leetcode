// 141. 环形链表 https://leetcode-cn.com/problems/linked-list-cycle/

// 给你一个链表的头节点 head ，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。如果链表中存在环 ，则返回 true 。 否则，返回 false 。

// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true

// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：true

// 示例 3：
// 输入：head = [1], pos = -1
// 输出：false

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表，时间复杂度 O(n)，空间复杂度 O(n)
// func hasCycle(head *ListNode) bool {
// 	seen := map[*ListNode]bool{}
// 	for head != nil {
// 		if seen[head] {
// 			return true
// 		}
// 		seen[head] = true
// 		head = head.Next
// 	}
// 	return false
// }

// 快慢指针，时间复杂度 O(n)，空间复杂度 O(1)
func hasCycle(head *ListNode) bool {
	// 如果链表为空或者只有一个节点，那么肯定没有环
	if head == nil || head.Next == nil {
		return false
	}

	// 快慢指针，从 head 之前的虚拟节点开始
	slow, fast := head, head.Next

	for slow != fast {
		// 如果快指针到达了链表尾部，那么肯定没有环
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}