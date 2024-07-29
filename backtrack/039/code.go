// 39. 组合总和 https://leetcode.cn/problems/combination-sum/

// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target
// 找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。

package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func([]int, int, int)
	dfs = func(comb []int, index, target int) {
		// 如果 target 为 0，说明找到了一个组合，将它放入结果中，然后返回
		if target == 0 {
			res = append(res, append([]int{}, comb...))
			return
		}

		// 从 index 开始遍历 candidates
		for i, c := range candidates[index:] {
			// 如果 c 小于等于 target，将 c 放入组合中，然后递归
			if c <= target {
				// 注意这里的 index+i，因为 candidates 中的数字可以重复使用，所以下一轮搜索的起点仍然是 index+i
				// target - c 为下一轮搜索的目标
				dfs(append(comb, c), index+i, target-c)
			}
		}
	}

	dfs(nil, 0, target)
	return res
}
