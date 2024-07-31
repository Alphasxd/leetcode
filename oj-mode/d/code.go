// 计算一系列数的和 https://ac.nowcoder.com/acm/problem/18207

// 输入描述:
// 输入数据包括多组。
// 每组数据一行,每行的第一个整数为整数的个数n(1 <= n <= 100), n为0的时候结束输入。
// 接下来n个正整数,即需要求和的每个正整数。
// 输出描述:
// 每组数据输出求和的结果

// 示例
// 输入
// 4 1 2 3 4
// 5 1 2 3 4 5
// 0
// 输出
// 10
// 15

package main

import "fmt"

func main() {
    var n int
    for {
        fmt.Scan(&n)
        if n == 0 {
            break
        }
        sum := 0
        for i := 0; i < n; i++ {
            var x int
            fmt.Scan(&x)
            sum += x
        }
        fmt.Println(sum)
    }
}
