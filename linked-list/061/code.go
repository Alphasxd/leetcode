// 61. 旋转链表 https://leetcode.cn/problems/rotate-list/

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

package leetcode

type ListNode struct {
    Val  int
    Next *ListNode
}

// 先将链表闭合成环，然后找到相应的位置断开这个环，确定新的链表头和链表尾
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }

    // 计算链表长度
    n := 1
    cur := head
    for cur.Next != nil {
        cur = cur.Next
        n++
    }

    // 将链表闭合成环
    cur.Next = head

    // 找到相应的位置断开这个环
    cur = head
    for i := 0; i < n-k%n-1; i++ {
        cur = cur.Next
    }

    // 新的链表头和链表尾
    newHead := cur.Next
    cur.Next = nil

    return newHead
}
