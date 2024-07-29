// 12. 整数转罗马数字 https://leetcode.cn/problems/integer-to-roman/

// 罗马数字包含以下七种字符；I，V，X，L，C，D 和 M。分别对应的数值为；1，5，10，50，100，500，1000。
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，这个特殊的规则只适用于以下：
// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

package leetcode

func intToRoman(num int) string {
	// 罗马数字中小的数字在大的数字的右边，从大到小排列
	romanTable := []struct {
		value int
		str   string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	// 存储结果
	var ans []byte
	// 遍历罗马数字表
	for _, v := range romanTable {
		// 如果当前数值比 num 小，则将 num 减去当前数值，同时将当前数值对应的罗马数字添加到结果中
		for num >= v.value {
			num -= v.value
			ans = append(ans, v.str...)
		}
		// 如果 num 为 0，则跳出循环
		if num == 0 {
			break
		}
	}
	return string(ans)
}
