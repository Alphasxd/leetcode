// 3. 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1：
// 输入：s = "abcabcbb"
// 输出：3

// 示例 2：
// 输入：s = "bbbbb"
// 输出：1

// 示例 3：
// 输入：s = "pwwkew"
// 输出：3

package leetcode

// func lengthOfLongestSubstring(s string) int {

//     // map[byte]int 用于存储字符最后出现的位置
//     index := make(map[byte]int)

//     length, i := 0, 0
//     // i 为子串起始位置，j 为子串结束位置
//     for j, c := range []byte(s) {
//         // 判断字符是否出现过，更新子串起始位置
//         if index[c] > i {
//             // 由于index[c]已经加1，所以可以直接赋值给i
//             i = index[c]
//         }
//         // 更新子串长度
//         if j-i+1 > length {
//             length = j - i + 1
//         }
//         // 更新字符最后出现的位置，下次遇到相同字符时，可以直接跳过
//         index[c] = j + 1
//     }
//     return length
// }

// 优化：使用rune代替byte，支持中文
// 思路：滑动窗口，使用map存储字符最后出现的位置，left为子串起始位置，right为子串结束位置
// 当遇到重复字符时，更新left的值，left的值为重复字符的下一个位置
func lengthOfLongestSubstring(s string) int {
    // map[rune]int 用于存储字符最后出现的位置
    m := make(map[rune]int)
    length, left := 0, 0
    // range s 会将字符串转换为rune数组
    for right, c := range s {
        // 判断字符是否出现过，更新子串起始位置
        if _, ok := m[c]; ok && m[c] >= left {
            // 字符已经出现过，更新left的值
            left = m[c] + 1
        }
        // 更新字符最后出现的位置
        m[c] = right
        if right-left+1 > length {
            length = right - left + 1
        }
    }
    return length
}
