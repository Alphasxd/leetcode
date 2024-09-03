// 26. 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度，元素的相对顺序应保持一致

// 示例 1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2]

// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]

package leetcode

func removeDuplicates(nums []int) int {
    n := len(nums)
    // nums 元素个数小于 2 时，直接返回原数组长度
    if n < 2 {
        return n
    }
    slow := 1// 快慢指针
    for fast := 1; fast < n; fast++ {
        if nums[slow-1] != nums[fast] {
            nums[slow] = nums[fast]
            slow++
        }
    }
    return slow
}
