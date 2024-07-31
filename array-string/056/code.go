// 56. 合并区间 https://leetcode.cn/problems/merge-intervals/

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

// 示例1：
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]

// 示例2：
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]

package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
    // 按每个区间的左端点升序排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res := [][]int{}
    prev := intervals[0]

    // 遍历区间，从第二个区间开始
    for i := 1; i < len(intervals); i++ {
        // 区间不重叠，直接将区间添加到结果集
        if cur := intervals[i]; prev[1] < cur[0] {
            res = append(res, prev)
            prev = cur
        } else { // 两个区间存在重叠，合并区间，新区间右端点取两个区间右端点的最大值
            prev[1] = max(prev[1], cur[1])
        }
    }
    // 将最后一个区间添加到结果集
    res = append(res, prev)
    return res
}
