// 79. 单词搜索 https://leetcode.cn/problems/word-search/

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。
// 如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
// 同一个单元格内的字母不允许被重复使用。

// 示例:
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true

package leetcode

func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    // 用于标记是否已经访问过
    used := make([][]bool, m)
    for i := range used {
        used[i] = make([]bool, n)
    }

    var canFind func(r, c, i int) bool
    canFind = func(r, c, i int) bool {
        // i == len(word) 说明已经找到了
        if i == len(word) {
            return true
        }

        // 越界
        if r < 0 || r >= m || c < 0 || c >= n {
            return false
        }

        // 已经访问过或者不符合当前字符
        if used[r][c] || board[r][c] != word[i] {
            return false
        }

        // 标记为已访问
        used[r][c] = true

        // 间接实现了回溯和剪枝
        if canFind(r-1, c, i+1) || canFind(r+1, c, i+1) || canFind(r, c-1, i+1) || canFind(r, c+1, i+1) {
            return true
        } else {
            used[r][c] = false // 重新标记为未访问, 因为下一次可能会访问到
            return false
        }
    }

    // 遍历所有的起点
    for i := range board {
        for j := range board[i] {
            if canFind(i, j, 0) {
                return true
            }
        }
    }
    return false
}
