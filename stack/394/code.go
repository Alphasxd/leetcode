// 394. 字符串解码 https://leetcode.cn/problems/decode-string/

// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// 示例 1：
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"

// 示例 2：
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"

package leetcode

import "strings"

// 思路：
// 1.使用两个栈，一个用于保存重复次数，一个用于保存当前处理的字符串
// 2.遍历字符串：
// - 如果是数字，读取完整的数字
// - 当遇到左括号`[`，说明要开始一个新的，完整的编码字符串，将当前的重复次数和字符串分别入压入数字栈和字符串栈，并重置重复次数和字符串
// - 当遇到字母时，直接添加到当前的字符串中，等待下一步处理
// - 当遇到右括号`]`，说明当前的编码字符串结束，需要将当前的字符串重复指定的次数，然后与上一个字符串合并
// 3.遍历结束后，返回字符串栈中的字符串
func decodeString(s string) string {
	// 数字栈和字符串栈
	cntStack, strStack := []int{}, []string{}
	// 当前的重复次数和字符串
	currNum, currStr := 0, ""

	// 遍历字符串
	for _, v := range s {
		switch {
		case v >= '0' && v <= '9':
			currNum = currNum*10 + int(v-'0')
		case v == '[':
			cntStack = append(cntStack, currNum) // 将当前的重复次数压入数字栈
			strStack = append(strStack, currStr) // 将当前的字符串压入字符串栈
			currNum, currStr = 0, ""             // 重置当前的重复次数和字符串
		case v == ']':
			// 从数字栈中弹出一个重复次数，从字符串栈中弹出一个字符串
			num, str := cntStack[len(cntStack)-1], strStack[len(strStack)-1]
			cntStack, strStack = cntStack[:len(cntStack)-1], strStack[:len(strStack)-1]
			// 将当前的字符串重复指定的次数，然后与上一个字符串合并
			currStr = str + strings.Repeat(currStr, num)
		default:
			currStr += string(v)
		}
	}
	return currStr
}
