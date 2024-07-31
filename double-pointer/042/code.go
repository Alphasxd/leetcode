// 42. 接雨水 https://leetcode.cn/problems/trapping-rain-water/

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6

// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9

package leetcode

func trap(height []int) int {
    left, right, preMax, sufMax, ans := 0, len(height)-1, 0, 0, 0
    // 使用双指针，从两边向中间遍历
    for left < right {
        // 前缀最大值 = max(前缀最大值, height[left])
        preMax = max(preMax, height[left])
        // 后缀最大值 = max(后缀最大值, height[right])
        sufMax = max(sufMax, height[right])
        // 如果 preMax < sufMax，说明左边的最大值小于右边的最大值，此时可以确定左边的最大值，计算左边的雨水量
        if preMax < sufMax {
            ans += preMax - height[left]
            left++
        } else {
            ans += sufMax - height[right]
            right--
        }
    }
    return ans
}
