// 36.有效的数独 https://leetcode-cn.com/problems/valid-sudoku/

// 请你判断一个 9 x 9 的数独是否有效。
// 数字 1-9 在每一个大的 3x3 的格子里只能出现一次。

package leetcode

func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]bool // 行、列、小九宫格
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// '.' 表示空格，跳过
			if board[i][j] == '.' {
				continue
			}
			cell := board[i][j] - '1' // 数字 1-9 转换为 0-8，正好对应下标
			// 分别判断行、列、小九宫格是否已经存在该数字
			// 对于小九宫格而言，(i/3, j/3) 表示在所在小九宫格里的坐标
			// 关键的一点是要把小九宫格的坐标转换为一维数组的下标，所以横坐标要乘以 3，即 i/3*3+j/3
			if row[i][cell] || col[j][cell] || box[i/3*3+j/3][cell] {
				return false
			}
			row[i][cell] = true
			col[j][cell] = true
			box[i/3*3+j/3][cell] = true
		}
	}
	return true
}