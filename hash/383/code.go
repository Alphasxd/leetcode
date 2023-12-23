// 383. 赎金信 https://leetcode-cn.com/problems/ransom-note

// 给你两个字符串 ransomNote 和 magazine ，判断ransomNote能否由 magazine 里面的字符构成。
// 如果可以构成，返回 true ；否则返回 false 。
// magazine 中的每个字符只能用一次。ransomNote和 magazine都只包含小写字母。

// 示例 1：
// 输入：ransomNote = "a", magazine = "b"
// 输出：false

// 示例 2：
// 输入：ransomNote = "aa", magazine = "ab"
// 输出：false

// 示例 3：
// 输入：ransomNote = "aa", magazine = "aab"
// 输出：true

package leetcode

// 只需要 magazine中的每个字符的出现次数都大于等于ransomNote中的字符出现次数即可，和242题有效的字母异位词类似
func canConstruct(ransomNote string, magazine string) bool {
	// 如果ransomNote的长度大于magazine，直接返回false
	if len(ransomNote) > len(magazine) {
		return false
	}
	// 用数组代替哈希表，因为字符的范围是a-z，所以数组长度为26
	cnt := [26]int{}
	// 统计magazine中每个字符出现的次数
	for _, ch := range magazine {
		// ch - 'a' 将字母表映射到数组下标 [0-25]
		cnt[ch-'a']++
	}
	// 遍历ransomNote，减去magazine中出现的字符
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		// 如果出现次数小于0，说明ransomNote中的字符在magazine中没有匹配的
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}