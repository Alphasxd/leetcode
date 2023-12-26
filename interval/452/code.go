// 452. 用最少数量的箭引爆气球 https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/

package leetcode

import "sort"

func findMinArrowShots(points [][]int) int {

	// 按照区间的右端点排序，这样可以尽可能少的使用箭
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// 至少有一个区间不相交，只需要记录不相交区间的数量即可
	ans := 1
	// 第一个区间的右端点
	prev := points[0][1]
	// 遍历所有区间，如果当前区间的左端点 > prev，说明不相交
	for _, p := range points {
		if p[0] > prev {
			ans++
			// 更新 prev
			prev = p[1]
		}
	}

	return ans
}
