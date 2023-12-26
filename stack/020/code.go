// 20.有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效
// 有效字符串需满足：
// 1.左括号必须用相同类型的右括号闭合
// 2.左括号必须以正确的顺序闭合
// 3.每个右括号都有一个对应的相同类型的左括号

package leetcode

func isValid(s string) bool {
	// 创建一个括号对应表
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 用栈存储左括号
	stack := []byte{}
	for _, ch := range []byte(s) {
		// 如果当前字符是右括号且右括号是括号对应表中的键
		if pair, ok := pairs[ch]; ok {
			// 如果栈为空或者栈顶元素不是当前字符对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			// 匹配成功，弹出栈顶元素，即对应的左括号
			stack = stack[:len(stack)-1]
		} else {
			// 如果当前字符是左括号，入栈
			stack = append(stack, ch)
		}
	}
	// 如果栈为空，说明所有左括号都被匹配
	return len(stack) == 0
}