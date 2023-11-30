// 24. 两两交换链表中的节点 https://leetcode-cn.com/problems/swap-nodes-in-pairs/

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]

// 示例 2：
// 输入：head = []
// 输出：[]

// 示例 3：
// 输入：head = [1]
// 输出：[1]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归，时间复杂度O(n)，空间复杂度O(n)
// func swapPairs(head *ListNode) *ListNode {
// 	// 如果链表为空或者只有一个节点，直接返回
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	// 最后返回的新链表的头节点
// 	newHead := head.Next

// 	// 递归调用，传入的参数是下一次递归的头节点，head指向下一次递归的头节点
// 	head.Next = swapPairs(newHead.Next)
// 	// newHead指向head
// 	newHead.Next = head

// 	// 返回新链表的头节点
// 	return newHead
// }

// 迭代，时间复杂度O(n)，空间复杂度O(1)
func swapPairs(head *ListNode) *ListNode {
	// 创建一个哑节点（虚拟头结点），指向 head
	dummy := &ListNode{Next: head}
	// 从哑节点开始迭代
	temp := dummy
	// 迭代条件：temp.Next 和 temp.Next.Next 都不为空
	for temp.Next != nil && temp.Next.Next != nil {
		first := temp.Next
		second := temp.Next.Next
		// 交换 first 和 second
		temp.Next = second
		first.Next = second.Next
		second.Next = first
		// first成为下一轮待交换的前置节点
		temp = first
	}
	// 返回哑节点的下一个节点
	return dummy.Next
}
