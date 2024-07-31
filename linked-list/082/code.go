// 82. 删除排序链表中的重复元素 II https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

// 给定一个已排序的链表的头 head，删除原始链表中所有重复数字的节点，只保留原始链表中没有重复出现的数字。返回同样按升序排列的结果链表。

package leetcode

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    // 创建虚拟头结点
    dummy := &ListNode{Next: head}
    curr := dummy
    for curr.Next != nil && curr.Next.Next != nil {
        if curr.Next.Val == curr.Next.Next.Val {
            x := curr.Next.Val
            // 删除所有重复的元素
            for curr.Next != nil && curr.Next.Val == x {
                curr.Next = curr.Next.Next
            }
        } else {
            curr = curr.Next
        }
    }
    return dummy.Next
}
