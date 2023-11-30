// 27. 移除元素 https://leetcode-cn.com/problems/remove-element/

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
// func removeElement(nums []int, val int) int {
// 	// 当作 index 和 返回值
// 	var i int
// 	// 遍历数组，如果元素不等于 val，就把该元素放到数组的 i 位置
// 	for _, num := range nums {
// 		if num != val {
// 			nums[i] = num
// 			i++
// 		}
// 	}
// 	return i
// }

// 双指针优化
func removeElement(nums []int, val int) int {
	// 双指针分别指向数组的头和尾
	left, right := 0, len(nums)
	for left < right {
		// 如果左指针指向的元素等于 val，就把右指针指向的元素放到左指针指向的位置
		// 如果移动过后，左指针指向的元素还是等于 val，就继续将右指针指向的元素放到左指针指向的位置，直到左指针指向的元素不等于 val
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}