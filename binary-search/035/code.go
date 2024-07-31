// 35. 搜索插入位置 https://leetcode.cn/problems/search-insert-position/

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2

// 示例 2:
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1

// 示例 3:
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4

package leetcode

func searchInsert(nums []int, target int) int {
    i, j := 0, len(nums)
    // 二分查找
    for i < j {
        mid := int(uint(i+j) >> 1)
        switch {
        case nums[mid] < target:
            i = mid + 1
        case nums[mid] > target:
            j = mid
        default:
            return mid
        }
    }
    // 未找到，返回插入位置，即 i，是指向第一个大于等于 target 的元素
    return i
}
