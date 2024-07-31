// 86. 分隔链表 https://leetcode.cn/problems/partition-list/

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，
// 使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。保留原始相对位置。

package leetcode

type ListNode struct {
    Val  int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    // 新建两个链表，一个存储小于x的节点，一个存储大于等于x的节点
    less, greater := new(ListNode), new(ListNode)
    // 两个游标指向两个链表的头节点
    lessPtr, greaterPtr := less, greater
    for head != nil {
        if head.Val < x {
            lessPtr.Next = head
            lessPtr = lessPtr.Next
        } else {
            greaterPtr.Next = head
            greaterPtr = greaterPtr.Next
        }
        head = head.Next
    }
    // 将大于等于x的链表的尾节点指向nil
    greaterPtr.Next = nil
    // 将小于x的链表的尾节点指向大于等于x的链表的头节点
    lessPtr.Next = greater.Next
    // 返回小于x的链表的头节点
    return less.Next
}
