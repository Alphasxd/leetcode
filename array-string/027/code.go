// 27. 移除元素 https://leetcode.cn/problems/remove-element/

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 示例 1：
// 输入：nums = [3,2,2,3], val = 3
// 输出：2, nums = [2,2]

// 示例 2：
// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
// 输出：5, nums = [0,1,4,0,3]

package leetcode

// 双指针
func removeElement(nums []int, val int) int {
    slow, fast := 0, 0 // 快慢指针
    for fast < len(nums) {
        // 只将不等于 val 的元素放到数组的 slow 位置
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    // slow已经自增了，所以直接返回 slow 即可
    return slow
}
