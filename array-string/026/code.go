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
    // nums 元素个数小于 2 时，直接返回原数组长度
    if len(nums) < 2 {
        return len(nums)
    }
    slow, fast := 0, 1 // 快慢指针
    for fast < len(nums) {
        if nums[slow] != nums[fast] {
            slow++
            nums[slow] = nums[fast]
        }
        fast++
    }
    // slow 为下标，长度为下标 + 1
    return slow + 1
}
