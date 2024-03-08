// 67. 二进制求和 https://leetcode-cn.com/problems/add-binary/

// 给你两个二进制字符串，返回它们的和（用二进制表示）。

// 示例1:
// 输入: a = "11", b = "1"
// 输出: "100"

package leetcode

func addBinary(a, b string) string {
	var ans []byte
	var carry byte
	// 从后往前遍历，直到两个字符串都遍历完
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		// sum 为当前位的和
		sum := carry
		// 如果 a 还有位数，sum 加上 a 的当前位
		if i >= 0 {
			sum += a[i] - '0'
		}
		// 如果 b 还有位数，sum 加上 b 的当前位
		if j >= 0 {
			sum += b[j] - '0'
		}
		// sum % 2 为当前位的值，sum / 2 为进位
		ans = append(ans, sum%2+'0')
		carry = sum / 2
	}
	// 如果还有进位，加上进位
	if carry > 0 {
		ans = append(ans, '1')
	}
	// 反转 ans
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}