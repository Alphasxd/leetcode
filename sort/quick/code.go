package leetcode

func quickSort(nums []int, lo, hi int) {
	if lo < hi {
		pivot := partition(nums, lo, hi)
		quickSort(nums, lo, pivot-1)
		quickSort(nums, pivot+1, hi)
	}
}

// partition函数用于对数组nums的子区间[lo, hi]进行一次划分
// 返回值是枢轴的最终位置
func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	i, j := lo, hi
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
}
