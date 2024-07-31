// 84. 柱状图中最大的矩形

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

// 示例 1:
// 输入: [2,1,5,6,2,3]
// 输出: 10

// 示例 2:
// 输入: [2,4]
// 输出: 4

package leetcode

// 暴力解法
// func largestRectangleArea(heights []int) int {
//     var maxArea int
//     for i, h := range heights {
//         var left, right = i, i
//         for left >= 0 && heights[left] >= h {
//             left--
//         }
//         for right < len(heights) && heights[right] >= h {
//             right++
//         }
//         if area := h * (right - left - 1); area > maxArea {
//             maxArea = area
//         }
//     }
//     return maxArea
// }

// 单调栈解法
func largestRectangleArea(heights []int) int {
    maxArea, stack := 0, []int{}
    heights = append(heights, 0)
    for i, h := range heights {
        // 如果当前高度小于栈顶高度，说明栈顶高度的右边界找到了，开始计算栈顶高度的面积
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
            // 弹出栈顶高度
            var height = heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            // 计算宽度
            var width int
            // 如果栈为空，说明栈顶高度左边界找不到，宽度为当前索引
            if len(stack) == 0 {
                width = i
            } else {
                // 否则宽度为当前索引减去栈顶索引减一
                width = i - stack[len(stack)-1] - 1
            }
            if area := height * width; area > maxArea {
                maxArea = area
            }
        }
        stack = append(stack, i)
    }
    return maxArea
}
