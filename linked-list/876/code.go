// 876. 链表的中间结点 https://leetcode.cn/problems/middle-of-the-linked-list/

// 给你单链表的头节点 head ，请你找出并返回中间节点。如果有两个中间节点，返回第二个中间节点。

package leetcode

type ListNode struct {
    Val  int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

// 判断链表是否有环，仅需在上面的基础上判断 slow == fast 即可
// func hasCycle(head *ListNode) bool {
//     slow, fast := head, head

//     for fast != nil && fast.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//         if slow == fast {
//             return true
//         }
//     }
//     return false
// }
