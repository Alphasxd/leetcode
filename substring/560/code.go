// 560. 和为 K 的子数组 https://leetcode.cn/problems/subarray-sum-equals-k/

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2

// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2

package leetcode

func subarraySum(nums []int, k int) int {
    cnt, preSum := 0, 0
    // 定义一个map，key为前缀和，value为前缀和出现的次数
    m := make(map[int]int)
    m[0] = 1 // 前缀和初始化
    for i := 0; i < len(nums); i++ {
        preSum += nums[i]
        if _, ok := m[preSum-k]; ok {
            cnt += m[preSum-k] // 将前缀和为preSum-k的个数，加到cnt中
        }
        m[preSum]++ // 更新前缀和个数
    }
    return cnt
}
