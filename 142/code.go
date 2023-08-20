// 142. 环形链表 II https://leetcode-cn.com/problems/linked-list-cycle-ii/

// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
// 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 不允许修改 链表。

// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点

// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点

// 示例 3：
// 输入：head = [1], pos = -1
// 输出：返回 null

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表，时间复杂度 O(n)，空间复杂度 O(n)
// func detectCycle(head *ListNode) *ListNode {
// 	seen := map[*ListNode]bool{}
// 	for head != nil {
// 		if seen[head] {
// 			return head
// 		}
// 		seen[head] = true
// 		head = head.Next
// 	}
// 	return nil
// }

// 快慢指针，时间复杂度 O(n)，空间复杂度 O(1)
// 设链表中环外部分的长度为 a，slow 指针进入环后，又走了 x 的距离与 fast 相遇，环长为 r
// 有：fast = 2 * slow => a + nr + x = 2 * (a + x) => a = nr - x
// 所以当 slow 和 fast 在环内相遇时，再同时从 head 和 相遇点 slow 出发，再次相遇时即为环的入口
// 从 head 出发走的距离为 a，从相遇点 slow 出发走的距离为 nr - x

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
