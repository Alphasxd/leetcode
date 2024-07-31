// 239. 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。

// 示例 1：
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]

// 示例 2：
// 输入：nums = [1], k = 1
// 输出：[1]

package leetcode

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 {
        return nums
    }
    var res []int
    // len(nums)-(k-1) 为滑动窗口的移动次数
    for i := 0; i < len(nums)-k+1; i++ {
        // nums[i:i+k] 为滑动窗口切片
        res = append(res, max(nums[i:i+k]))
    }
    return res
}

func max(nums []int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}
