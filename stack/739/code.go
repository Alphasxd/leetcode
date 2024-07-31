// 739. 每日温度 https://leetcode.cn/problems/daily-temperatures/

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 示例 1:
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]

// 示例 2:
// 输入: temperatures = [30,40,50,60]
// 输出: [1,1,1,0]

// 示例 3:
// 输入: temperatures = [30,60,90]
// 输出: [1,1,0]

package leetcode

// 思路：使用单调栈解决，栈中存储的是数组的下标，栈中的元素从栈底到栈顶对应的温度依次递减。
func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures)) // 结果切片
    stack := []int{}                      // 单调栈
    for i, v := range temperatures {
        for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
            // 当前温度大于栈顶元素对应的温度时，栈顶元素出栈，计算结果
            res[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i) // 当前元素下标入栈
    }
    return res
}
