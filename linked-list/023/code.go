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

func mergeKLists(lists []*ListNode) *ListNode {
    pq := make(PriorityQueue, 0)
    // 初始化堆
    heap.Init(&pq)
    // 将所有链表的头结点加入堆
    for _, head := range lists {
        if head != nil {
            heap.Push(&pq, head)
        }
    }
    dummy := new(ListNode)
    curr := dummy
    // 从堆中取出最小的节点，加入到合并链表中
    for pq.Len() > 0 {
        // 取出最小的节点
        min := heap.Pop(&pq).(*ListNode)
        // 将最小节点的下一个节点加入堆
        if min.Next != nil {
            heap.Push(&pq, min.Next)
        }
        // 将最小节点加入到合并链表中
        curr.Next = min
        curr = curr.Next
    }
    return dummy.Next
}

// 实现一个最小堆
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
    *pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() any {
    old := *pq
    n := len(old)
    x := old[n-1]
    *pq = old[:n-1]
    return x
}
