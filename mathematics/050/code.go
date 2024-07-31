// 50. 求 x 的 n 次幂函数
// 自己实现 pow(x, n) ，即计算 x 的 n 次幂函数

package leetcode

// 快速幂
func myPow(x float64, n int) float64 {
    // 如果 n 为负数，那么相当于求 (1/x)^(-n)
    if n < 0 {
        x = 1 / x
        n = -n
    }
    ans := 1.0
    for n > 0 {
        // 将 n 转换为二进制数，从低位到高位遍历
        // 如果当前位为 1，那么需要乘上将 ans 乘上 x
        if n&1 == 1 {
            ans *= x
        }
        // 按位平方，每一位都是上一位的平方
        x *= x
        // n 右移一位
        n >>= 1
    }
    return ans
}
