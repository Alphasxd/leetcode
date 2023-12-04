// 13. 罗马数字转整数 https://leetcode-cn.com/problems/roman-to-integer/

// 罗马数字包含以下七种字符；I，V，X，L，C，D 和 M。分别对应的数值为；1，5，10，50，100，500，1000。
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，这个特殊的规则只适用于以下：
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

package leetcode

var romanTable = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

// 思路：从左到右遍历字符串，如果当前字符对应的数值比后一个字符对应的数值小，
// 则减去当前字符对应的数值，否则加上当前字符对应的数值
func romanToInt(s string) int {
	length := len(s)
	// 存储结果，初始化为第一个字符对应的数值
	ans := romanTable[s[0]]
	// 待遍历字符的前一个字符对应的数值，初始化为第一个字符对应的数值
	pre := ans
	// 从第二个字符开始遍历
	for i := 1; i < length; i++ {
		// 当前字符对应的数值
		cur := romanTable[s[i]]
		// 加上当前字符对应的数值
		ans += cur
		// 如果当前字符对应的数值比前一个字符对应的数值大，
		// 则减去前一个字符对应的数值的两倍，首先减去前一个字符对应的数值，再减去当前字符对应的数值
		if pre < cur {
			ans -= pre * 2
		}
		pre = cur
	}
	return ans
}
