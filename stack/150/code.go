// 150. 逆波兰表达式求值 https://leetcode.cn/problems/evaluate-reverse-polish-notation/

package leetcode

import "strconv"

// 后缀表达式求值
// 1. 遇到数字，入栈
// 2. 遇到运算符，弹出栈顶的两个元素，运算后将结果入栈
// 3. 最后栈中只剩一个元素，即为结果
func evalRPN(tokens []string) int {
    stack := make([]int, 0)
    for _, token := range tokens {
        switch token {
        case "+", "-", "*", "/":
            num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
            stack = stack[:len(stack)-2]
            stack = append(stack, calc(num1, num2, token))
        default:
            // 将字符串转换为数字
            num, _ := strconv.Atoi(token)
            stack = append(stack, num)
        }
    }
    return stack[0]
}
func calc(num1, num2 int, op string) int {
    switch op {
    case "+":
        return num1 + num2
    case "-":
        return num1 - num2
    case "*":
        return num1 * num2
    default:
        return num1 / num2
    }
}
