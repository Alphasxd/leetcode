// 057. 插入区间 https://leetcode.cn/problems/insert-interval/

// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然 有序且不重叠 （如果有必要的话，可以 合并区间 ）。

package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	ans := make([][]int, 0)
	var placed bool

	for _, interval := range intervals {
		// 在插入区间的右侧且无交集，说明后续的区间不再有交集
		if interval[0] > right {
			// 且未插入新区间，此时插入新区间并标记已插入
			if !placed {
				ans = append(ans, []int{left, right})
				placed = true
			}
			// 将该区间加入答案数组
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集，直接将该区间加入答案数组
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	// 预防插入区间在最后一个位置
	if !placed {
		ans = append(ans, []int{left, right})
	}

	return ans
}
