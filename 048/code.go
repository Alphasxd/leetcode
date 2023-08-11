// 48. 旋转图像 https://leetcode-cn.com/problems/rotate-image

// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]

// 示例 2：
// 输入：matrix = [[5, 1, 9,11],[2, 4, 8,10],[13, 3, 6, 7],[15,14,12,16]]
// 输出：[[15,13, 2, 5],[14, 3, 4, 1],[12, 6, 8, 9],[16, 7,10,11]]

package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)

	// 水平翻转
	for i := 0; i < n/2; i++ {
		// i 行和 n-i-1 行交换
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ { // 只需要遍历对角线上半部分
			// matrix[i][j] 和 matrix[j][i] 交换
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
