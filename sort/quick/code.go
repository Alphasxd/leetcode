package leetcode

func quickSort(nums []int, lo, hi int) {
	if lo < hi {
		// 通过partition函数找到枢轴的位置
		pivot := partition(nums, lo, hi)
		// 对枢轴左侧的子区间进行快速排序
		quickSort(nums, lo, pivot-1)
		// 对枢轴右侧的子区间进行快速排序
		quickSort(nums, pivot+1, hi)
	}
}

// partition函数用于对数组nums的子区间[lo, hi]进行一次划分
// 返回值是枢轴的最终位置
func partition(nums []int, lo, hi int) int {
	// 选择nums[lo]作为枢轴
	pivot := nums[lo]
	i, j := lo, hi
	// 循环直到i和j相遇
	for i < j {
		// 从右向左找到第一个小于枢轴的元素
		for i < j && nums[j] >= pivot {
			j--
		}
		// 将这个元素移动到左侧
		nums[i] = nums[j]
		// 从左向右找到第一个大于枢轴的元素
		for i < j && nums[i] <= pivot {
			i++
		}
		// 将这个元素移动到右侧
		nums[j] = nums[i]
	}
	// 将枢轴放到最终位置
	nums[i] = pivot
	// 返回枢轴的位置
	return i
}
