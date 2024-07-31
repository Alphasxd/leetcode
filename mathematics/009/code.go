// 9. 回文数 https://leetcode.cn/problems/palindrome-number/

package leetcode

func isPalindrome(x int) bool {
    // 负数不是回文数，末尾为 0 的数也不是回文数
    if x < 0 || (x > 0 && x%10 == 0) {
        return false
    }
    var y int
    // 将 x 的后半部分反转，与前半部分比较
    // 当 x <= y 时，反转已经完成
    for x > y {
        y = y*10 + x%10
        x /= 10
    }
    // x == y 为偶数位数
    // x == y/10 为奇数位数，此时只需要将 y 去掉最后一位即可
    return x == y || x == y/10
}
