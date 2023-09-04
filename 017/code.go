// 电话号码的字母组合

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。
// 1:!@#   2:abc   3:def
// 4:ghi   5:jkl   6:mno
// 7:pqrs  8:tuv   9:wxyz
// *: +    0: _    #: %

// 示例 1：
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

package leetcode

func letterCombinations(digits string) []string {
	// 如果字符串为空，直接返回空
	if len(digits) == 0 {
		return nil
	}

	buttons := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}


	var results []string
	// 因为有几个数字，组合就有几位，所以直接初始化
	temp := make([]byte, len(digits))

	var dfs func(int)
	dfs = func(i int) {
		// 如果已经遍历到最后一位，就把结果加入到结果集中
		if i == len(digits) {
			results = append(results, string(temp))
			return
		}

		// 获取当前数字对应字母的 byte 数组
		letters := buttons[digits[i]-'2']

		// 遍历字母数组，把每个字母加入到临时数组中
		for j := 0; j < len(letters); j++ {
			temp[i] = letters[j]
			dfs(i + 1)
		}
	}

	// 从参数 digits 的第一位开始遍历
	dfs(0)

	return results
}
