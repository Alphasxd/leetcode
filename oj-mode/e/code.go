// 输入描述:
// 输入的第一行包括一个正整数t(1 <= t <= 100), 表示数据组数。
// 接下来t行, 每行一组数据。
// 每行的第一个整数为整数的个数n(1 <= n <= 100)。
// 接下来n个正整数, 即需要求和的每个正整数。
// 输出描述:
// 每组数据输出求和的结果

// 输入
// 2
// 4 1 2 3 4
// 5 1 2 3 4 5
// 输出
// 10
// 15

package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			fmt.Println(0)
			continue
		}
		var ans int
		for j := 0; j < n; j++ {
			var num int
			fmt.Scan(&num)
			ans += num
		}
		fmt.Println(ans)
	}
}