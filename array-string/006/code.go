// 6. N字形变换 https://leetcode-cn.com/problems/zigzag-conversion/

// 将一个给定字符串s根据给定的行数numRows，以从上往下、从左到右进行Z字形排列。
// 比如输入字符串为"PAYPALISHIRING"行数为3时，排列如下：
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

package leetcode

func convert(s string, numRows int) string {
	// 如果只有一行，直接返回
	if numRows == 1 {
		return s
	}

	// 初始化二维数组，切片数组的数组长度为numRows，也就是有numRows行
	rows := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		// 每行的切片容量为len(s)/2+1，这样可以避免后面append时扩容
		rows[i] = make([]byte, 0, len(s)/2+1)
	}

	// 从第一行开始遍历，direction为1表示向下，-1表示向上
	curRow := 0
	direction := -1
	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		// 如果到了第一行或者最后一行，改变方向
		if curRow == 0 || curRow == numRows-1 {
			direction = -direction
		}
		// 根据方向，向上或向下移动一行
		curRow += direction
	}

	res := make([]byte, 0, len(s))
	for _, row := range rows {
		res = append(res, row...)
	}
	// 将切片转换为字符串
	return string(res)	
}