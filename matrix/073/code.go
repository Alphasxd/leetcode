// 73. 矩阵置零 https://leetcode-cn.com/problems/set-matrix-zeroes/

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

// 示例 1：
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]

// 示例 2：
// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

package leetcode

// 方法一：使用标记数组，时间复杂度 O(mn)，空间复杂度 O(m+n)
// func setZeroes(matrix [][]int)  {
// 	row := make([]bool, len(matrix))
// 	col := make([]bool, len(matrix[0]))

// 	for i, r := range matrix {
// 		for j, v := range r {
// 			if v == 0 {
// 				row[i] = true
// 				col[j] = true
// 			}
// 		}
// 	}

// 	for i, r := range matrix {
// 		for j := range r {
// 			if row[i] || col[j] {
// 				r[j] = 0
// 			}
// 		}
// 	}
// }

// 方法二：使用两个标记变量，时间复杂度 O(mn)，空间复杂度 O(1)
func setZeroes(matrix [][]int) {
	// 第一行和第一列的标记变量
	row0, col0 := false, false

	// 遍历第一行，判断第一行是否有 0
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}

	// 遍历第一列，判断第一列是否有 0
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}

	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				// 将第一行和第一列作为标记数组
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据第一行和第一列的标记数组，将矩阵的元素置零, 从第二行和第二列开始
	// 注意：第一行和第一列的标记数组不需要处理
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 根据 row0 和 col0，将第一行和第一列置零
	if row0 {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
