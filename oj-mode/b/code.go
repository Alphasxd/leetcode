// 计算a+b https://ac.nowcoder.com/acm/problem/18205

// 输入描述:
// 输入第一行包括一个数据组数t(1 <= t <= 100)
// 接下来每行包括两个正整数a,b(1 <= a, b <= 1000)
// 输出描述:
// 输出a+b的结果

// 示例
// 输入
// 2
// 1 5
// 10 20
// 输出
// 6
// 30

package main

import "fmt"

func main() {
    var t, a, b int
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        fmt.Scan(&a, &b)
        fmt.Println(a + b)
    }
}
