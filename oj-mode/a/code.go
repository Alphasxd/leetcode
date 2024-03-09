// 计算a+b https://ac.nowcoder.com/acm/contest/5657/A

// 输入描述：
// 输入包括两个正整数 a, b (1 <= a, b <= 10^9)，输入数据包括多组

// 输出描述：
// 输出a+b的结果

// 示例1
// 输入
// 1 5
// 10 20
// 输出
// 6
// 30

package main

import (
	"fmt"
	"io"
)

func main() {
	var a, b int
	for {
		if _, err := fmt.Scan(&a, &b); err != io.EOF {
			fmt.Println(a + b)
		} else {
			break
		}
	}
}