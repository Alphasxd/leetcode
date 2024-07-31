// 139. 单词拆分 https://leetcode.cn/problems/word-break/

// 给你一个字符串 s 和一个字符串列表 wordDict ，s 的所有字符都是小写英文字母。
// 如果可以利用 wordDict 中的单词拼接 s ，则返回 true ；否则，返回 false 。

// 示例 1：
// 输入：s = "leetcode", wordDict = ["leet", "code"]
// 输出：true

// 示例 2：
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false

package leetcode

// hash + dp
func wordBreak(s string, wordDict []string) bool {
    // 使用 map 存储 wordDict，方便查找
    dict := make(map[string]struct{})
    for _, word := range wordDict {
        dict[word] = struct{}{}
    }
    // ans[i] 表示 s[:i] 是否可以拆分成 wordDict 中的单词
    ans := make([]bool, len(s)+1)
    // 初始化 ans[0] 为 true
    ans[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            // 如果 ans[j] 为 true 且 s[j:i] 在 wordDict 中，则 ans[i] 为 true
            if _, ok := dict[s[j:i]]; ok && ans[j] {
                ans[i] = true
                break
            }
        }
    }
    return ans[len(s)]
}
