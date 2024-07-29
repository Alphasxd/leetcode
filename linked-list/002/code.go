// 2. 两数相加 https://leetcode.cn/problems/add-two-numbers/

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例 1：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

// 示例 2：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]

// 示例 3：
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
// 解释：9999999 + 9999 = 10009998.

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// 创建一个辅助节点，指向首元节点之前的节点，模拟头结点
	dummy := new(ListNode)
	// 创建一个游标
	cur := dummy

	// 存放进位值
	var carry int

	// 遍历两个链表，有一个不为空就继续遍历
	for l1 != nil || l2 != nil {
		// 创建新节点用于存放相加后的值，因为返回一个新链表，所以每次都要创建一个新节点
		cur.Next = new(ListNode)
		cur = cur.Next

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		// 将相加后的值存入新节点
		cur.Val = carry % 10
		// 计算进位值
		carry /= 10
	}

	// 如果最后还有进位值，就再创建一个新节点，此时carry不需要再除以10
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
