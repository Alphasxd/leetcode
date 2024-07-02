// 131. 分割回文串 https://leetcode-cn.com/problems/palindrome-partitioning/

// 给你一个字符串 s, 请你将 s 分割成一些子串, 使每个子串都是回文串. 返回 s 所有可能的分割方案.

package leetcode

func partition(s string) [][]string {
	path := []string{}
	ans := [][]string{}
	n := len(s)

	var dfs func(int, int)
	dfs = func(index, start int) {
		if index == n {
			ans = append(ans, append([]string(nil), path...)) // 拷贝一份 path
			return
		}

		// 不选 index 和 index+1 之间的逗号
		if index < n-1 {
			dfs(index+1, start)
		}

		// 选 index 和 index+1 之间的逗号, 把 s[index] 作为子串的最后一个字符
		if isPalindrome(s, start, index) {
			path = append(path, s[start:index+1])
			dfs(index+1, index+1) // 下一个子串的起始位置是 index+1
			path = path[:len(path)-1] // 恢复 path
		}
	}

	dfs(0, 0)
	return ans
}

// 判断是否是回文串
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
