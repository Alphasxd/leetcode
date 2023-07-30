// 15. 三数之和 https://leetcode-cn.com/problems/3sum/

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
// 同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]

// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]

// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]

package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	// 将 nums 排序
	sort.Ints(nums)

	//var solution [][]int
	solution := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			switch {
			// 如果 sum > 0，说明 nums[hi] 太大，hi--
			case sum > 0:
				hi--
			//	如果 sum < 0，说明 nums[lo] 太小，lo++
			case sum < 0:
				lo++
			//	如果 sum == 0，说明找到了一组解，lo++，hi--
			default:
				solution = append(solution, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--
				// 去重
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			}
		}
	}
	return solution
}
