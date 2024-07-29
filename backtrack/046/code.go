// 46. 全排列 https://leetcode.cn/problems/permutations/

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

package leetcode

func permute(nums []int) [][]int {
	// 定义一个二维切片，用于存储所有的排列结果
	var res [][]int
	// 定义一个递归函数，用于生成所有的排列结果
	var f func([]int, []int)
	f = func(nums, path []int) {
		// 如果 nums 中的所有数都已经被使用过了，将 path 添加到 res 中
		if len(nums) == 0 {
			res = append(res, path)
			return
		}
		// 从 nums 中取出一个数，放入 path 中，然后递归
		for i, v := range nums {
			// 将 nums[:i] 和 nums[i+1:] 拼接起来，得到一个新的切片，再将 v 添加到末尾
			// 这样就得到了一个新的切片，其中不包含原始切片 nums 中的第 i 个数
			newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
			// 将 v 添加到 path 中，得到一个新的切片
			newPath := append(path, v)
			// 递归调用 dfs 函数，生成所有的排列结果
			f(newNums, newPath)
		}
	}
	// 调用 dfs 函数，生成所有的排列结果
	f(nums, []int{})
	// 返回所有的排列结果
	return res
}
