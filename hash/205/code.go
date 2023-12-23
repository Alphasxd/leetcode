// 205. 同构字符串 https://leetcode-cn.com/problems/isomorphic-strings/

// 给定两个字符串 s 和 t，判断它们是否是同构的。
// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。

// 示例 1:
// 输入: s = "egg", t = "add"
// 输出: true

// 示例 2:
// 输入: s = "foo", t = "bar"
// 输出: false

package leetcode

// 双向映射，与290题单词规律类似
func isIsomorphic(s string, t string) bool {
	// 如果两个字符串长度不相等，直接返回 false
	if len(s) != len(t) {
		return false
	}
	// 需要两个 map 来保证双向映射
	m1, m2 := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		// 如果两个 map 中都没有记录，则添加记录，byte的默认值为0
		if m1[s[i]] == 0 && m2[t[i]] == 0 {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		} else if m1[s[i]] != t[i] || m2[t[i]] != s[i] {
			// 如果两个 map 中有记录，但是不对应，则返回 false
			return false
		}
	}
	return true
}