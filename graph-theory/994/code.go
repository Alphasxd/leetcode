// 994. 腐烂的橘子 https://leetcode.cn/problems/rotting-oranges/

// 在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

package leetcode

func orangesRotting(grid [][]int) int {
	cnts, cost, m, n, queue := 0, 0, len(grid), len(grid[0]), [][]int{}
	for i, r := range grid {
		for j, v := range r {
			if v > 0 {
				// 统计所有橘子的数量
				cnts++
				if v == 2 {
					// 将所有腐烂的橘子加入队列
					queue = append(queue, []int{i, j})
				}
			}
		}
	}
	for len(queue) > 0 {
		l := len(queue)
		// 减去腐烂的橘子，剩下的就是新鲜的橘子
		cnts -= l
		for i := 0; i < l; i++ {
			// 从队列中取出腐烂的橘子
			point := queue[0]
			// 更新队列
			queue = queue[1:]
			// 将腐烂的橘子周围的新鲜橘子腐烂
			for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				nx, ny := point[0]+d[0], point[1]+d[1]
				// 判断是否越界，以及是否为新鲜橘子
				if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
					// 将新鲜橘子腐烂并加入队列
					grid[nx][ny] = 2
					queue = append(queue, []int{nx, ny})
				}
			}
		}
		// 如果队列不为空，说明有新鲜橘子被腐烂
		if len(queue) > 0 {
			cost++
		}
	}
	// 如果还有新鲜橘子，返回 -1，否则返回腐烂的时间
	if cnts == 0 {
		return cost
	}
	return -1
}
