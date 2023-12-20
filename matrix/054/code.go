// 54. 螺旋矩阵 https://leetcode-cn.com/problems/spiral-matrix/

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]

// 示例 2：
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

package leetcode

// 思路：按照顺时针顺序遍历矩阵，每次遍历完一圈，收缩边界，继续遍历下一圈，直到遍历完所有元素
func spiralOrder(matrix [][]int) []int {
	// 结果数组
	res := []int{}

	if len(matrix) == 0 {
		return res
	}
	// 上下左右边界
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	for top <= bottom && left <= right {
		// 从左到右
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		// 从上到下，i从top+1开始
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		// 在保证有上下左右边界的情况下，进行下面两个循环
		if top < bottom && left < right {
			// 从右到左，j从right-1开始
			for j := right - 1; j > left; j-- {
				res = append(res, matrix[bottom][j])
			}
			// 从下到上
			for i := bottom; i > top; i-- {
				res = append(res, matrix[i][left])
			}
		}
		// 完成一圈，边界收缩
		left++
		right--
		top++
		bottom--
	}

	return res
}
