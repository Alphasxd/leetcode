// 最长递增子序列 https://leetcode.cn/problems/longest-increasing-subsequence/

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 示例 1：
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4

// 示例 2：
// 输入：nums = [0,1,0,3,2,3]
// 输出：4

package leetcode

func lengthOfLIS(nums []int) int {
    var res int
    // dp[i]表示以nums[i]结尾的最长递增子序列的长度
    dp := make([]int, len(nums))
    // 初始化dp数组，每个元素都至少可以单独成为一个子序列
    for i := range dp {
        dp[i] = 1
    }
    for i := range nums {
        for j := range i {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    for _, v := range dp {
        res = max(res, v)
    }
    return res
}
