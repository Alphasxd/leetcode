// 22. 括号生成 https://leetcode-cn.com/problems/generate-parentheses/

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

package leetcode

func generateParenthesis(n int) []string {
	pair := make([]byte, n*2)
	var dfs func([]string, []byte, int, int, int) []string
	dfs = func(pairs []string, pair []byte, n, left, right int) []string {
		// 如果左和右括号都用完了，就加入到结果中
		if left == n && right == n {
			return append(pairs, string(pair))
		}
		// 如果左括号还有剩余，就可以放一个左括号
		if left < n {
			pair[left+right] = '('
			pairs = dfs(pairs, pair, n, left+1, right)
		}
		// 如果右括号的数量小于左括号的数量，就可以放一个右括号
		if right < left {
			pair[left+right] = ')'
			pairs = dfs(pairs, pair, n, left, right+1)
		}
		return pairs
	}
	return dfs(nil, pair, n, 0, 0)
}