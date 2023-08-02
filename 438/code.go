// 438. 找到字符串中所有字母异位词 https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

// 示例 1:
// 输入: s: "cbaebabacd" p: "abc"
// 输出: [0, 6]

// 示例 2:
// 输入: s: "abab" p: "ab"
// 输出: [0, 1, 2]

package leetcode

func findAnagrams(s string, p string) []int {
	// 定义一个数组来存储结果
	var res []int
	// 定义一个数组来存储 p 中每个字符出现的次数
	var cnt [26]int
	// 遍历 p，统计每个字符出现的次数
	for _, c := range p {
		cnt[c-'a']++
	}
	// 定义滑动窗口的左右边界
	left, right := 0, 0
	// 定义一个数组来存储窗口中每个字符出现的次数
	var window [26]int
	// 遍历 s
	for right < len(s) {
		// 将当前字符出现的次数加 1
		window[s[right]-'a']++
		// 如果当前字符出现的次数大于 p 中该字符出现的次数
		// 则说明该字符在窗口中出现的次数过多，需要移动左边界
		for window[s[right]-'a'] > cnt[s[right]-'a'] {
			// 将左边界字符出现的次数减 1
			window[s[left]-'a']--
			// 左边界右移
			left++
		}
		// 如果窗口的大小与 p 的大小相等
		// 则说明找到了一个字母异位词
		if right-left+1 == len(p) {
			// 将左边界索引加入结果数组
			res = append(res, left)
		}
		// 右边界右移
		right++
	}
	return res
}
