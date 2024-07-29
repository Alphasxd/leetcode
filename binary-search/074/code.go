// 74. 搜索二维矩阵 https://leetcode.cn/problems/search-a-2d-matrix/

// 给你一个满足下述两条属性的 m x n 整数矩阵：
// 每行中的整数从左到右按非递减顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

package leetcode

import "sort"

// 从右上角开始搜索，如果当前值比 target 小，则向下搜索，如果当前值比 target 大，则向左搜索
// 时间复杂度：O(m+n)，空间复杂度：O(1)
// func searchMatrix(matrix [][]int, target int) bool {
// 	m, n := len(matrix), len(matrix[0])
// 	i, j := 0, n-1
// 	for {
// 		if matrix[i][j] == target {
// 			return true
// 		}
// 		if matrix[i][j] < target {
// 			i++
// 		} else {
// 			j--
// 		}
// 		if i >= m || j < 0 {
// 			break
// 		}
// 	}
// 	return false
// }

// 二分查找
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}
