// 160. 相交链表 https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

// 示例 1：
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
// 输出：Intersected at '8'

// 示例 2：
// 输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Intersected at '2'

// 示例 3：
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null

package leetcode

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// Map 时间复杂度O(m+n) 空间复杂度O(m)或O(n)
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
//     vis := map[*ListNode]bool{}
// 	for i := headA; i != nil; i = i.Next {
// 		vis[i] = true
// 	}

// 	for j := headB; j != nil; j = j.Next {
// 		// 如果存在相同节点，直接返回，此题不存在有相交节点后面不相交的情况
// 		if vis[j] {
// 			return j
// 		}
// 	}

// 	return nil
// }

// 双指针 时间复杂度O(m+n) 空间复杂度O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果有一个链表为空，直接返回
	if headA == nil || headB == nil {
		return nil
	}

	// 两个指针分别指向两个链表的头节点
	pA, pB := headA, headB

	// 如果两个指针不相等，继续循环
	for pA != pB {
		if pA == nil {
			// 如果pA为空，将pA指向headB
			pA = headB
		} else {
			// 否则，将pA指向pA的下一个节点
			pA = pA.Next
		}

		if pB == nil {
			// 如果pB为空，将pB指向headA
			pB = headA
		} else {
			// 否则，将pB指向pB的下一个节点
			pB = pB.Next
		}
	}

	// pA和pB相等时，跳出循环，返回pA或pB
	return pA
}