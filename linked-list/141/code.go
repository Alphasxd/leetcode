// 141. 环形链表 https://leetcode.cn/problems/linked-list-cycle/

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

// 集合，时间复杂度 O(n)，空间复杂度 O(n)
func hasCycle1(head *ListNode) bool {
    set := make(map[*ListNode]struct{})
    for head != nil {
        if _, ok := set[head]; ok {
            return true
        }
        set[head] = struct{}{}
        head = head.Next
    }
    return false
}

// 快慢指针，时间复杂度 O(n)，空间复杂度 O(1)
func hasCycle(head *ListNode) bool {
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
