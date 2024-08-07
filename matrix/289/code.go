// 289. 生命游戏 https://leetcode.cn/problems/game-of-life/

// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：
// 1 即为活细胞（live），或 0 即为死细胞（dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都
// 遵循以下四条生存定律：
// 1. 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 2. 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 3. 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 4. 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
// 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。
// 给你 m x n 网格面板 board 的当前状态，返回下一个状态。

package leetcode

func gameOfLife(board [][]int) {
    m, n := len(board), len(board[0])
    // 用以下数字分别代表细胞的 4 种状态
    // 0: dead -> dead
    // 1: live -> live
    // 2: live -> dead
    // 3: dead -> live
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            // 统计遍历到的细胞的周围活细胞数量
            live := 0
            // x的作用是遍历当前细胞的上下细胞，同时保证不越界
            // y的作用是遍历当前细胞的左右细胞，同时保证不越界
            // 忽略掉当前细胞本身，没有必要统计
            for x := max(i-1, 0); x <= min(i+1, m-1); x++ {
                for y := max(j-1, 0); y <= min(j+1, n-1); y++ {
                    if x == i && y == j {
                        continue
                    }
                    // 1, 2 的当前状态都是 live
                    if board[x][y] == 1 || board[x][y] == 2 {
                        live++
                    }
                }
            }
            // 根据规则更新当前细胞状态，涉及到状态变化的只有1，3，4三种情况
            // 忽略第 2 条规则，因为不涉及状态变化，
            // 和题目中没有给出的规则 dead -> dead 一样，不需要处理
            if board[i][j] == 1 && (live < 2 || live > 3) {
                // 符合 1, 3 规则，都是 live -> dead，用 2 标记
                board[i][j] = 2
            }
            if board[i][j] == 0 && live == 3 {
                // 符合 4 规则，是 dead -> live，用 3 标记
                board[i][j] = 3
            }
        }
    }
    // 将board中的数字mod2，即可得到下一状态
    // 0，1, 2, 3 -> 0, 1, 0, 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            board[i][j] %= 2
        }
    }
}
