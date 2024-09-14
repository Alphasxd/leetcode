// 206.反转链表 https://leetcode.cn/problems/reverse-linked-list/

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

// 直接调转指针，时间复杂度O(n)，空间复杂度O(1)
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

// 头插法，时间复杂度O(n)，空间复杂度O(1)
func reverseList1(head *ListNode) *ListNode {
    // 创建哑节点
    dummy := &ListNode{}
    // 依次遍历原链表，头插法插入到 dummy 后面，最后返回 dummy.Next 即可
    for head != nil {
        // 辅助节点，防止断链
        next := head.Next
        // 头插法
        head.Next = dummy.Next
        dummy.Next = head
        head = next
    }
    return dummy.Next
}
