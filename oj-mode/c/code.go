// 计算a+b https://ac.nowcoder.com/acm/problem/18206

// 输入描述:
// 输入包括两个正整数a,b(1 <= a, b <= 10^9),输入数据有多组, 如果输入为0 0则结束输入
// 输出描述:
// 输出a+b的结果

// 示例
// 输入
// 1 5
// 10 20
// 0 0
// 输出
// 6
// 30

package main

import "fmt"

func main() {
	var a, b int
	for {
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}