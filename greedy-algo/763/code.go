// 763. 划分字母区间 https://leetcode.cn/problems/partition-labels/

// 给你一个字符串 s，请你对 s 的子串进行划分，并返回子串的个数。
// 划分要尽可能多的子串，并且每个字母只能出现在一个子串中。将划分结果按顺序连接，得到的字符串应当与原字符串完全相同。

// 示例 1：
// 输入：s = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca", "defegde", "hijhklij"

package leetcode

func partitionLabels(s string) []int {
    // 记录每个字符最后出现的位置，通过遍历实现更新
    lastIndex := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        lastIndex[s[i]] = i
    }

    var ans []int
    start, end := 0, 0
    // 遍历字符串，end 记录当前子串的结束位置
    for i := 0; i < len(s); i++ {
        // 更新 end 为当前字符的最后出现位置
        end = max(end, lastIndex[s[i]])
        // 当前位置等于 end 时，表示当前子串结束，可以划分
        if i == end {
            ans = append(ans, end-start+1)
            // 更新 start 为下一个子串的开始位置
            start = end + 1
        }
    }
    return ans
}
