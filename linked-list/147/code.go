// 147. 对链表进行插入排序 https://leetcode-cn.com/problems/insertion-sort-list/

// 给定单个链表的头 head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。
// 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
// 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
// 重复直到所有输入数据插入完为止。

// 示例 1：
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]

// 示例 2：
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	// 创建哑节点，指针域指向head
	dummy := &ListNode{Next: head}

	// 从 head 的第二个节点开始遍历
	curr := head.Next

	// 断开 head 与后面的节点的连接
	head.Next = nil

	for curr != nil {
		// 防止断链
		next := curr.Next
		// 每次都需要从表头开始寻找插入位置
		pre := dummy
		// 遍历已排序的链表，找到插入位置，即比 curr.Val 大的节点
		for pre.Next != nil && pre.Next.Val < curr.Val {
			pre = pre.Next
		}
		curr.Next = pre.Next
		pre.Next = curr
		curr = next
	}

	return dummy.Next
}
