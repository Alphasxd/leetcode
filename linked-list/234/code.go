// 234. 回文链表 https://leetcode.cn/problems/palindrome-linked-list/

// 给你一个单链表的头节点 head (事实上是首元结点)，请你判断该链表是否为回文链表。如果是，返回 true；否则，返回 false。

// 示例 1：
// 输入：head = [1,2,2,1]
// 输出：true

// 示例 2：
// 输入：head = [1,2]
// 输出：false

package leetcode

type ListNode struct {
    Val  int
    Next *ListNode
}

// 切片存储，时间复杂度：O(n)，空间复杂度：O(n)
// func isPalindrome(head *ListNode) bool {
//     // 创建一个切片
//     vals := []int{}
//     // 遍历链表，将链表的值存入切片
//     for head != nil {
//         vals = append(vals, head.Val)
//         head = head.Next
//     }
//     // 遍历切片的前半部分，如果有不相等的值，返回false
//     for i, v := range vals[:len(vals)/2] {
//         if v != vals[len(vals)-i-1] {
//             return false
//         }
//     }
//     // 遍历完切片，没有返回false，说明是回文链表，返回true
//     return true
// }

// 快慢指针，时间复杂度：O(n)，空间复杂度：O(1)
func isPalindrome(head *ListNode) bool {
    var prev *ListNode
    slow, fast := head, head
    // 快指针走到链表尾部，慢指针走到链表中间
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        // 在遍历过程中，将前半部分链表反转
        next := slow.Next
        slow.Next = prev
        prev = slow
        slow = next
    }

    // head 指向前半部分已经反转的首元结点
    head = prev

    // prev 再指回slow，用于再次恢复前半部分的链表
    prev = slow

    // 如果 fast 不为 nil 说明跳出循环的时候 fast.Next 为 nil，即链表长度为奇数
    // 将 slow 跳过中间节点，指向后半部分的首元结点
    if fast != nil {
        slow = slow.Next
    }
    // 用于判断是否是回文链表
    palindrome := true
    // 从中间向两边遍历，判断是否是回文链表
    for head != nil {
        if head.Val != slow.Val {
            palindrome = false
        }
        // 再将前半部分反转回来
        next := head.Next
        head.Next = prev
        prev = head
        head = next // 同时将 head 往后移动（前半部分）

        slow = slow.Next // slow 往后移动（后半部分）
    }
    // 跳出循环，说明是回文链表，返回 true
    return palindrome
}
