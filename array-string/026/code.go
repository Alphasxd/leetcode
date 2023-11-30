// 26. 删除有序数组中的重复项 https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度，元素的相对顺序应保持一致

// 示例 1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2]

// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]

package leetcode

func removeDuplicates(nums []int) int {
	// 数组长度为0，直接返回0
	if len(nums) == 0 {
		return 0
	}

	// 定义两个指针，一个指针i用于遍历数组，另一个指针j用于记录不重复元素的位置
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			// 如果nums[j] != nums[i]，说明nums[i]是不重复的元素，将nums[i]赋值给nums[j+1]
			j++
			nums[j] = nums[i]
		}
	}
	// 下标j从0开始，所以返回j+1
	return j + 1
}