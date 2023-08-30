// 200. 岛屿数量 https://leetcode-cn.com/problems/number-of-islands/

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

// 示例 1：
// 输入：grid = [
// 	["1","1","1","1","0"],
// 	["1","1","0","1","0"],
// 	["1","1","0","0","0"],
// 	["0","0","0","0","0"]
// ]
// 输出：1

// 示例 2：
// 输入：grid = [
// 	["1","1","0","0","0"],
// 	["1","1","0","0","0"],
// 	["0","0","1","0","0"],
// 	["0","0","0","1","1"]
// ]
// 输出：3

package leetcode

func numIslands(grid [][]byte) int {
	// 初始化 visited 二维数组
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	// 岛屿数量
	var num int
	for i, r := range grid {
		for j, c := range r {
			if c == '0' || visited[i][j] {
				continue
			}
			num++
			visit(grid, visited, i, j)
		}
	}
	return num
}

// 深度优先搜索
func visit(grid [][]byte, visited [][]bool, i, j int) {
	// 如果是水或者已经访问过了，就不再访问
	if grid[i][j] == '0' || visited[i][j] {
		return
	}
	// 标记为已访问
	visited[i][j] = true
	// 访问上下左右，将其临近的陆地都标记为已访问，成为一个岛屿
	if i > 0 {
		visit(grid, visited, i-1, j)
	}
	if i < len(grid)-1 {
		visit(grid, visited, i+1, j)
	}
	if j > 0 {
		visit(grid, visited, i, j-1)
	}
	if j < len(grid[i])-1 {
		visit(grid, visited, i, j+1)
	}
}
