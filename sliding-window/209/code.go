// 209. 长度最小的子数组 https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其和 ≥ target 的长度最小的连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
// 并返回其长度。如果不存在符合条件的子数组，返回 0 。

// 示例 1：
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2

// 示例 2：
// 输入：target = 4, nums = [1,4,4]
// 输出：1

package leetcode

import (
	"math"
	"sort"
)

// brute force
// func minSubArrayLen(target int, nums []int) int {
// 	n := len(nums)
// 	if n == 0 {
// 		return 0
// 	}
// 	ans := math.MaxInt32
// 	for i := 0; i < n; i++ {
// 		sum := 0;
// 		for j := i; j < n; j++ {
// 			sum += nums[j]
// 			if sum >= target {
// 				ans = min(ans, j-i+1)
// 				break
// 			}
// 		}
// 	}
// 	if ans == math.MaxInt32 {
// 		return 0
// 	}
// 	return ans
// }

// 前缀和 + 二分查找
func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    ans := math.MaxInt32
	// sums[i] 为 nums[0..i-1] 的和
    sums := make([]int, n+1)
    for i := 1; i <= n; i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }
    for i := 1; i <= n; i++ {
        targetSum := target + sums[i-1]
		// 返回第一个大于等于 targetSum 的元素的索引，如果不存在这样的元素，返回 len(sums):n+1
        bound := sort.SearchInts(sums, targetSum)
        if bound != n+1 {
			// bound-(i-1) 即为满足条件的子数组的长度
            ans = min(ans, bound-(i-1))
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }
    return ans
}