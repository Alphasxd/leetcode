// 205. 同构字符串 https://leetcode.cn/problems/isomorphic-strings/

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
func isIsomorphic(s, t string) bool {
    m1 := map[byte]byte{}
    m2 := map[byte]byte{}
    for i := range s {
        x, y := s[i], t[i]
        if m1[x] > 0 && m1[x] != y || m2[y] > 0 && m2[y] != x {
            return false
        }
        m1[x] = y
        m2[y] = x
    }
    return true
}
