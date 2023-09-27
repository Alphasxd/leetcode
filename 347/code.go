// 347. 前 K 个高频元素 https://leetcode-cn.com/problems/top-k-frequent-elements/

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现`频率`前 k 高的元素。可以按 任意顺序 返回答案。

// 示例 1:
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]

// 示例 2:
// 输入: nums = [1], k = 1
// 输出: [1]

package leetcode

import (
	"container/heap"
	// "sort"
)

// 方法一：使用 map 统计元素出现的频率，然后对 map 中的元素按照出现的频率进行排序，最后返回前 k 个元素
// 时间复杂度：O(nlogn)，其中 n 为数组的长度
// func topKFrequent(nums []int, k int) []int {
// 	ans := []int{}
// 	// 创建一个 map，用于统计每个元素出现的频率，key 为元素值，value 为元素出现的频率
// 	freqMap := map[int]int{}
// 	for _, v := range nums {
// 		freqMap[v]++
// 	}
// 	// 将 map 中的元素添加到 ans 中
// 	for k, _ := range freqMap {
// 		ans = append(ans, k)
// 	}

// 	// 通过 sort.Slice() 函数对 ans 进行排序，排序规则为：按照元素出现的频率降序排序
// 	sort.Slice(ans, func(i, j int) bool {
// 		return freqMap[ans[i]] > freqMap[ans[j]]
// 	})

// 	// 返回 ans 的前 k 个元素，即为前 k 个高频元素
// 	return ans[:k]
// }

// 方法二：使用小顶堆，堆中存储的是元素值，堆中元素的个数为 k，堆中元素按照元素出现的频率进行排序
// 时间复杂度：O(nlogk)，其中 n 为数组的长度
func topKFrequent(nums []int, k int) []int {
	hp := NumFreqHeap{}
	// heapify
	heap.Init(&hp)

	// 创建一个 map，用于统计每个元素出现的频率，key 为元素值，value 为元素出现的频率
	freqMap := map[int]int{}
	for _, v := range nums {
		freqMap[v]++
	}

	// 将 map 中的元素添加到 hp 中
	for kk, vv := range freqMap {
		// 如果 hp 中元素的个数小于 k，则直接将元素添加到 hp 中
		if hp.Len() < k {
			heap.Push(&hp, &NumFreq{Num: kk, Freq: vv})
		} else {
			// 如果 hp 中元素的个数等于 k，则将元素与堆顶元素进行比较
			// 如果元素的频率大于堆顶元素的频率，则将堆顶元素删除，将元素添加到 hp 中
			if vv > hp[0].Freq {
				heap.Pop(&hp)
				heap.Push(&hp, &NumFreq{kk, vv})
			}
		}
	}

	// 将 hp 中的元素添加到 ans 中
	ans := []int{}
	for hp.Len() > 0 {
		ans = append(ans, heap.Pop(&hp).(*NumFreq).Num)
	}

	return ans
}

// NumFreq 用于存储元素值和元素出现的频率
type NumFreq struct {
	Num  int // 元素值
	Freq int // 元素出现的频率
}

// NumFreqHeap 用于存储 NumFreq，堆中元素按照 NumFreq.Freq 进行排序
type NumFreqHeap []*NumFreq

// 实现小顶堆，让频率大的元素下沉，频率小的元素上浮方便 Pop
func (h NumFreqHeap) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h NumFreqHeap) Len() int {
	return len(h)
}

func (h NumFreqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 入堆
func (h *NumFreqHeap) Push(x any) {
	*h = append(*h, x.(*NumFreq))
}

// 出堆
func (h *NumFreqHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}