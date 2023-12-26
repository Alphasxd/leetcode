package leetcode

import "strings"

func simplifyPath(path string) string {
	res := []string{}
	// 将字符串以 / 分割成数组，存在空字符串，.，..，其他字符串
	for _, s := range strings.Split(path, "/") {
		// 如果是 .. 则删除上一级目录，即弹出栈顶元素
		if s == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else if s != "." && s != "" {
			// 否则如果不是 . 也不是空字符串，即为目录名，入栈
			res = append(res, s)
		}
	}
	// Unix绝对路径以 / 开头，并用 / 分隔目录名
	return "/" + strings.Join(res, "/")
}