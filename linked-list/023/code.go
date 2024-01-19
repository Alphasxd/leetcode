// 23. 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]

// 示例 2：
// 输入：lists = []
// 输出：[]

// 示例 3：
// 输入：lists = [[]]
// 输出：[]

package leetcode

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	hp := ListNodeHeap{}
	// 将每个待合并的链表的首元节点添加到小顶堆中
	for _, head := range lists {
		if head != nil {
			hp = append(hp, head)
		}
	}
	// 将 hp 初始化为堆，因为 hp 实现了 heap.Interface 接口，所以可以直接调用 heap.Init(hp)
	heap.Init(&hp)

	// 创建哑节点，用于返回合并后的链表
	dummy := new(ListNode)
	curr := dummy
	for hp.Len() > 0 {
		// 取出堆顶元素（最小值）
		min := heap.Pop(&hp).(*ListNode)
		// 如果最小值存在后继节点，则将后继节点添加到堆中
		if min.Next != nil {
			heap.Push(&hp, min.Next)
		}

		// 将最小值添加到合并后的链表中
		curr.Next = min
		curr = curr.Next

	}

	return dummy.Next

}

// 实现 heap.Interface 接口
// type Interface interface {
// 	sort.Interface
// Len() int
// Less(i, j int) bool
// Swap(i, j int)
// 	Push(x interface{}) // add x as element Len()
// 	Pop() interface{}   // remove and return element Len() - 1.
// }
func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}