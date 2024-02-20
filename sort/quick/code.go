package leetcode

func quickSort(nums []int, lo, hi int) {
	if lo < hi {
		pivot := partition(nums, lo, hi)
		quickSort(nums, lo, pivot-1)
		quickSort(nums, pivot+1, hi)
	}
}

// 一次划分，枢轴的位置是确定的，后续不再变化
// 枢轴左边的元素都小于枢轴，右边的元素都大于枢轴，但左右两边的元素不一定有序
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
