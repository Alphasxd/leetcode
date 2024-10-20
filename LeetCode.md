# Based on the Top 100 Liked Questions

### 哈希

<details markdown="1">
<summary>1. 两数之和（Easy）</summary>

> 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

> 可以假设每种输入只会对应一个答案，并且不能使用两次相同的元素。可以按任意顺序返回答案

```golang
func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for i, num := range nums {
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}
```
> 遍历原始数组将**数组元素**作为 key，**元素索引**作为 value 存入 map 即可；

> Tips：把返回结果写在新增 kv 之前

</details>

<details markdown="1">
<summary>49. 字母异位词分组（Medium）</summary>

> 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表

```golang
func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    for _, str := range strs {
        s := []byte(str)
        sort.Slice(s, func(i, j int) bool {
            return s[i] < s[j]
        })
        sortedStr := string(s)
        m[sortedStr] = append(m[sortedStr], str)
    }
    // 转为二维切片
    ans := make([][]string, 0, len(m))
    for _, v := range m {
        ans = append(ans, v)
    }
    return ans
}
```

> 字符串 `->` 字符切片 `->` 排序 `->` 字符串 `->` map

</details>

<details markdown="1">
<summary>128. 最长连续序列（Medium）</summary>

> 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

> 要求时间复杂度为 O(n)

```golang
func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    for _, num := range nums {
        m[num] = true
    }
    longest := 0
    for num := range m {
        // 确保 num 是连续序列的起点
        if m[num-1] {
            continue
        }
        length := 1
        for m[num+1] {
            num++
            length++
        }
        if length > longest {
            longest = length
        }
    }
    return longest
}
```

> 用 map 当作集合，遍历数组寻找到“连续序列的起点”，然后更新最长连续序列长度

</details>

### 双指针

<details markdown="1">
<summary>283. 移动零（Easy）</summary>

> 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序

> 必须在不复制数组的情况下原地对数组进行操作

```golang
func moveZeroes(nums []int)  {
    for i, j := 0, 0; j < len(nums); j++ {
        if nums[j] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
}
```

> i, j 都从 0 开始，j 负责遍历数组，i 负责记录非 0 元素的位置，j 遇到非 0 元素就交换 i 和 j 的值

</details>

<details markdown="1">
<summary>88. 合并两个有序数组（Easy）</summary>

> 给你两个非递减顺序排列的整数数组 nums1 和 nums2，有两个整数 m 和 n，分别表示 nums1 和 nums2 中的元素数目。请你将 nums2 合并到 nums1 中，使合并后的数组同样按非递减顺序排列。

> nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

```golang
func merge(nums1 []int, m int, nums2 []int, n int) {
    p1, p2, p := m-1, n-1, m+n-1
    for p2 >= 0 { // nums2 还有元素未处理
        if p1 >= 0 && nums1[p1] > nums2[p2] {
            nums1[p] = nums1[p1]
            p1--
        } else {
            nums1[p] = nums2[p2]
            p2--
        }
        p--
    }
}
```

> 倒序双指针，p1 指向 nums1 末尾，p2 指向 nums2 末尾，p 指向 nums1 的末尾，每次比较 p1 和 p2 的值，将较大的值放到 nums1 的末尾，避免覆盖 nums1 的值

</details>

<details markdown="1">
<summary>11. 盛水最多的容器（Medium）</summary>

> 给定一个长度为 n 的整数数组 height。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。

> 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。返回容器可以储存的最大水量。

```golang
func maxArea(height []int) int {
    ans := 0
    i, j := 0, len(height)-1
    for i < j {
        if height[i] < height[j] {
            ans = max(ans, height[i]*(j-i))
            i++
        } else {
            ans = max(ans, height[j]*(j-i))
            j--
        }
    }
    return ans
}
```
> 双指针，i 指向首，j 指向尾，每次移动 height 值较小的指针，计算面积并更新最大面积 `Area = min(height[i], height[j]) * (j - i)`

> 如果移动较大的那个指针，不会有任何意义，因为木桶原理，容器的容量只会有变得更小的可能

</details>

<details markdown="1">
<summary>15. 三数之和（Medium）</summary>

> 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

> 答案中不可以包含重复的三元组

```golang
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    ans := make([][]int, 0)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        lo, hi := i+1, len(nums)-1
        for lo < hi {
            sum := nums[i] + nums[lo] + nums[hi]
            switch {
            case sum > 0:
                hi--
            case sum < 0:
                lo++
            default:
                ans = append(ans, []int{nums[i], nums[lo], nums[hi]})
                hi--
                lo++
                // 去重，不能有重复的三元组
                for lo < hi && nums[lo] == nums[lo-1] {
                    lo++
                }
                for lo < hi && nums[hi] == nums[hi+1] {
                    hi--
                }
            }
        }
    }
    return ans
}
```

> 先排序，然后固定一个数，双指针遍历剩余数组，注意去重

</details>

<details markdown="1">
<summary>75. 颜色分类（Medium）</summary>

> 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。必须在不使用库内置的 sort 函数的情况下解决这个问题

```golang
func sortColors(nums []int) {
    zero, two := 0, len(nums)-1
    // one 从左往右遍历，遇到 0 则与 zero 交换，遇到 2 则与 two 交换
    for one := 0; one <= two; {
        switch nums[one] {
        case 0:
            nums[zero], nums[one] = nums[one], nums[zero]
            zero++
            one++
        case 1:
            one++
        case 2:
            nums[one], nums[two] = nums[two], nums[one]
            two--
        }
    }
}
```

> 三指针，zero 指向 0 的最右边界，two 指向 2 的最左边界，one 遍历数组，遇到 0 和 zero 交换，遇到 2 和 two 交换

</details>

<details markdown="1">
<summary>31. 下一个排列（Medium）</summary>

> 给你一个整数数组 nums ，找出 nums 的下一个排列。如果不存在下一个更大的排列，则将 nums 重新排列成最小的排列（即升序排列）

```golang
func nextPermutation(nums []int) {
    if len(nums) < 2 {
        return
    }
    i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
    for i >= 0 && nums[i] >= nums[j] {
        i--
        j--
    }
    if i >= 0 {
        for nums[i] >= nums[k] {
            k--
        }
        nums[i], nums[k] = nums[k], nums[i]
    }
    reverse(nums, j, len(nums)-1)
}
func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}
```
> 可以直接将字典排序理解为一个整数比较大小的问题，找到下一个比当前大的最小的整数即可

> 从右往左找到第一个升序对 (i, j)，再从右往左找到第一个大于 nums[i] 的数，交换 i 和 k，最后翻转 j 到末尾的元素

</details>

<details markdown="1">
<summary>287. 寻找重复数（Medium）</summary>

> 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。假设 nums 只有 一个重复的整数 ，返回 这个重复的数

> 必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间

```golang
func findDuplicate(nums []int) int {
    slow := nums[nums[0]]
    fast := nums[nums[nums[0]]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    // slow 从 nums[0] 开始，fast 从相遇点开始，再次相遇时即为重复元素
    duplicate := nums[0]
    for duplicate != slow {
        duplicate = nums[duplicate]
        slow = nums[slow]
    }
    return duplicate
}
```

> 快慢指针，类似于判断链表是否有环，找到环的入口

</details>

### 滑动窗口

<details markdown="1">
<summary>3. 无重复字符的最长子串（Medium）</summary>

> 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度

```golang
func lengthOfLongestSubstring(s string) int {
    m := make(map[rune]int)
    length, left := 0, 0
    // range s 会将字符串转换为 rune 数组
    for right, c := range s {
        if _, ok := m[c]; ok && m[c] >= left {
            // 字符已经出现过，更新left的值
            left = m[c] + 1
        }
        // 更新字符最后出现的位置
        m[c] = right
        if right-left+1 > length {
            length = right - left + 1
        }
    }
    return length
}
```

> 用 map 记录字符出现的位置，left 为子串起始，right 为子串结束，遇到重复字符时更新左指针 left

</details>

<details markdown="1">
<summary>438. 找到字符串中所有字母异位词（Medium）</summary>

> 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。异位词 指由相同字母重排列形成的字符串（包括相同的字符串）

```golang
func findAnagrams(s string, p string) []int {
    var res []int
    var cnt [26]int
    for _, c := range p {
        cnt[c-'a']++
    }
    left, right := 0, 0
    var window [26]int
    for right < len(s) {
        window[s[right]-'a']++
        for window[s[right]-'a'] > cnt[s[right]-'a'] {
            window[s[left]-'a']--
            left++
        }
        if right-left+1 == len(p) {
            res = append(res, left)
        }
        right++
    }
    return res
}
```

> 用数组记录 p 中字符出现次数，遍历 s，每次移动窗口，判断窗口内字符出现次数是否和 p 相同

</details>

### 子串

<details markdown="1">
<summary>560. 和为 K 的子数组（Medium）</summary>

> 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。子数组是数组中元素的连续非空序列。

```golang
func subarraySum(nums []int, k int) int {
    cnt, preSum := 0, 0
    // key 为前缀和，value 为前缀和出现的次数
    m := make(map[int]int)
    m[0] = 1
    for i := 0; i < len(nums); i++ {
        preSum += nums[i]
        if _, ok := m[preSum-k]; ok {
            cnt += m[preSum-k]
        }
        m[preSum]++
    }
    return cnt
}
```

> 用 map 记录前缀和出现的次数，遍历数组，计算前缀和，判断是否存在 preSum - k 的前缀和

</details>

### 字符串

<details markdown="1">
<summary>415. 字符串相加（Easy）</summary>

> 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

```golang
func addStrings(num1, num2 string) string {
    res := ""
    i, j, carry := len(num1)-1, len(num2)-1, 0
    for i >= 0 || j >= 0 {
        var n1, n2 int
        if i >= 0 {
            n1 = int(num1[i] - '0')
        }
        if j >= 0 {
            n2 = int(num2[j] - '0')
        }
        tmp := n1 + n2 + carry
        carry = tmp / 10
        res = strconv.Itoa(tmp%10) + res
        i--
        j--
    }
    if carry > 0 {
        res = "1" + res
    }
    return res
}
```

> 用 sum 记录当前和，max 记录最大和，遍历数组，如果 sum 小于 0，就从当前元素开始重新计算

</details>

<details markdown="1">
<summary>43. 字符串相乘（Medium）</summary>

> 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

```golang
func multiply(num1, num2 string) string {
    ints := make([]int, len(num1)+len(num2))
    // 数组 ints 保存乘积的每一位
    for i := 0; i < len(num1); i++ {
        for j := 0; j < len(num2); j++ {
            ints[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
        }
    }
    // 进位
    for i := len(ints) - 1; i > 0; i-- {
        ints[i-1] += ints[i] / 10
        ints[i] %= 10
    }
    return intsToString(ints)
}

// 将数组转换为字符串
func intsToString(ints []int) string {
    var i int
    for i < len(ints)-1 && ints[i] == 0 {
        i++
    }
    ints = ints[i:]

    var b strings.Builder
    b.Grow(len(ints))
    for _, i := range ints {
        b.WriteByte('0' + byte(i))
    }
    return b.String()
}
```

> 用数组保存乘积的每一位，最后将数组转换为字符串

</details>

<details markdown="1">
<summary>93. 复原 IP 地址（Medium）</summary>

> 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

```golang
func restoreIpAddresses(s string) []string {
	bytes := make([]string, 0, 4)
	return dfs(nil, bytes, s)
}
func dfs(IPs, bytes []string, s string) []string {
	// bytes 存储构建 IP 地址的每一段
	if len(bytes) == 4 {
		// bytes 中已经有 4 段，且 s 为空，说明找到了一个合法的 IP 地址
		if len(s) == 0 {
			IPs = append(IPs, strings.Join(bytes, "."))
		}
		return IPs
	}
	// 检查 s 是否以字符 '0' 开头，如果是，则只能将 '0' 作为一段加入 bytes
	if len(s) > 0 && s[0] == '0' {
		return dfs(IPs, append(bytes, "0"), s[1:])
	}
	// 遍历 s 的每一位，将其作为一段加入 bytes
	var num int
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			break
		}
		IPs = dfs(IPs, append(bytes, s[:i+1]), s[i+1:])
	}
	return IPs
}
```

> dfs 遍历 s 的每一位，将其作为一段加入 bytes，如果 bytes 中已经有 4 段，且 s 为空，说明找到了一个合法的 IP 地址

</details>

### 普通数组

<details markdown="1">
<summary>53. 最大子数组和（Medium）</summary>

> 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。子数组是数组中的一个连续部分

```golang
func maxSubArray(nums []int) int {
    var max, sum int
    for i, num := range nums {
        sum += num
        // 如果 sum 大于 max 或者是第一次遍历时
        if sum > max || i == 0 {
            max = sum
        }
        // 舍弃过去，重新开始，贪心思想
        if sum < 0 {
            sum = 0
        }
    }
    return max
}
```

> 用 sum 记录当前和，max 记录最大和，遍历数组，如果 sum 小于 0，就从当前元素开始重新计算

</details>

<details markdown="1">
<summary>152. 乘积最大子数组（Medium）</summary>

> 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

```golang
func maxProduct(nums []int) int {
    product, res := 1, nums[0]
    for i := range len(nums) {
        product *= nums[i]
        res = max(res, product)
        if nums[i] == 0 {
            // 重置 product，即当前子数组的乘积
            product = 1
        }
    }
    // 两次遍历，第一次从左到右，第二次从右到左
    product = 1
    for i := len(nums) - 1; i >= 0; i-- {
        product *= nums[i]
        res = max(res, product)
        if nums[i] == 0 {
            product = 1
        }
    }
    return res
}
```

> 两次遍历，第一次从左往右计算乘积，第二次从右往左计算乘积，取最大值，应对有奇数个负数的情况

</details>

<details markdown="1">
<summary>56. 合并区间（Medium）</summary>

> 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

```golang
func merge(intervals [][]int) [][]int {
    // 按每个区间的左端点升序排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res := [][]int{}
    prev := intervals[0]
    // 遍历区间，从第二个区间开始
    for i := 1; i < len(intervals); i++ {
        if cur := intervals[i]; prev[1] < cur[0] {
            res = append(res, prev)
            prev = cur
        } else { // 有重叠，更新 prev 的右端点为两个区间的最大值
            prev[1] = max(prev[1], cur[1])
        }
    }
    res = append(res, prev)
    return res
}
```

> 先排序，然后遍历数组，如果当前区间的左边界大于前一个区间的右边界，就添加到结果集，否则更新前一个区间的右边界（取最大值）

</details>

<details markdown="1">
<summary>189. 轮转数组（Medium）</summary>

> 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

```golang
func rotate(nums []int, k int) {
    // 对 k 取余，避免 k 大于数组长度时多余的反转
    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}
func reverse(nums []int) {
    // 只需要反转一半的数组，j-i-1 是为了避免奇数数组时中间的数被反转两次
    for i, n := 0, len(nums); i < n/2; i++ {
        nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
    }
}
```

> 三次翻转，先整体翻转，再翻转前 k 个元素，最后翻转后 n-k 个元素

</details>

<details markdown="1">
<summary>238. 除自身以外数组的乘积（Medium）</summary>

> 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

```golang
func productExceptSelf(nums []int) []int {
    length := len(nums)
    ans := make([]int, length)
    product := 1
    for i := range ans {
        ans[i] = product
        // 将 ans[i] 设为左侧所有元素的乘积
        product *= nums[i]
    }
    product = 1
    for i := length - 1; i >= 0; i-- {
        // 将上一步的结果乘以右侧所有元素的乘积即为答案
        ans[i] *= product
        // 将 ans[i] 设为 ans[i] * 右侧所有元素的乘积
        product *= nums[i]
    }
    return ans
}
```

> 左侧乘积 * 右侧乘积

</details>

<details markdown="1">
<summary>41. 缺失的第一个正数（Hard）</summary>

> 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

`集合`

```golang
func firstMissingPositive1(nums []int) int {
    set := make(map[int]struct{}, len(nums))
    for _, v := range nums {
        if v > 0 {
            set[v] = struct{}{}
        }
    }
    for i := 1; i <= len(nums); i++ {
        if _, ok := set[i]; !ok {
            return i
        }
    }
    return len(nums) + 1
}
```

> 遍历数组，将正数存入集合，再遍历 1 到 n+1，找到第一个不在集合中的正数

`原地哈希`

```golang
func firstMissingPositive(nums []int) int {
    for _, v := range nums {
        for v > 0 && v <= len(nums) && nums[v-1] != v {
            nums[v-1], v = v, nums[v-1]
        }
    }
    for i := 1; i <= len(nums); i++ {
        if nums[i-1] != i {
            return i
        }
    }
    return len(nums) + 1
}
```

> 自定义哈希，将每个正数放到对应的位置，再遍历数组，找到第一个不在对应位置的正数

</details>

### 矩阵

<details markdown="1">
<summary>48. 旋转图像（Medium）</summary>

> 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

```golang
func rotate(matrix [][]int) {
    n := len(matrix)
    // 水平翻转
    for i := 0; i < n/2; i++ {
        matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
    }
    // 主对角线翻转
    for i := 0; i < n; i++ {
        for j := 0; j < i; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
}
```

> 顺时针旋转90度： 转置 + 水平翻转

> 逆时针旋转90度： 转置 + 垂直翻转

</details>

<details markdown="1">
<summary>240. 搜索二维矩阵 II（Medium）</summary>

> 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：每行的元素从左到右升序排列。每列的元素从上到下升序排列。

```golang
func searchMatrix(matrix [][]int, target int) bool {
    for _, r := range matrix {
        i := sort.SearchInts(r, target)
        if i < len(r) && r[i] == target {
            return true
        }
    }
    return false
}
```

> sort.SearchInts() 返回 target 在 r 中的索引，如果找到了就返回 true

</details>

### 链表

<details markdown="1">
<summary>160. 相交链表（Easy）</summary>

> 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 

```golang
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pA, pB := headA, headB
    for pA != pB {
        if pA == nil {
            pA = headB
        } else {
            pA = pA.Next
        }
        if pB == nil {
            pB = headA
        } else {
            pB = pB.Next
        }
    }
    return pA
}
```

> 两个指针分别遍历两个链表，当遍历到尾部时，指向另一个链表的头部，最终会在相交节点相遇

</details>

<details markdown="1">
<summary>206. 反转链表（Easy）</summary>

> 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表

`头插法`

```golang
func reverseList(head *ListNode) *ListNode {
    dummy := &ListNode{}
    for head != nil {
        // 辅助节点，防止断链
        next := head.Next
        // 头插法
        head.Next = dummy.Next
        dummy.Next = head
        head = next
    }
    return dummy.Next
}
```

`反转指针`
```golang
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next // next 指向 curr 的下一个节点
        head.Next = prev  // 反转 curr 的指针
        prev = head       // prev 指向 curr
        head = next       // curr 指向 next
    }
    // prev 指向反转后的链表的头节点
    return prev
}
```

</details>

<details markdown="1">
<summary>143. 重排链表（Medium）</summary>

> 给定一个单链表 L 的头节点 head ，单链表 L 表示为： L0 → L1 → … → Ln - 1 → Ln

> 请将其重新排列后变为：L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → … 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。


```golang
func reorderList(head *ListNode) {
    // 找到链表中点
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 反转后半部分链表
    var prev *ListNode
    for slow != nil {
        next := slow.Next
        slow.Next = prev
        prev, slow = slow, next
    }
    // 合并两个链表
    for p, q := head, prev; p != q; p, q = q, p {
        next := p.Next
        p.Next = q
        p = next
    }
}
```
</details>

<details markdown="1">
<summary>86. 分隔链表（Medium）</summary>

> 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。保留原始相对位置。


```golang
func partition(head *ListNode, x int) *ListNode {
    // 新建两个链表，一个存储小于 x 的节点，一个存储大于等于 x 的节点
    less, greater := new(ListNode), new(ListNode)
    // 两个游标指向两个链表的头节点
    lessPtr, greaterPtr := less, greater
    for head != nil {
        if head.Val < x {
            lessPtr.Next = head
            lessPtr = lessPtr.Next
        } else {
            greaterPtr.Next = head
            greaterPtr = greaterPtr.Next
        }
        head = head.Next
    }
    // 将大于等于 x 的链表的尾节点指向 nil
    greaterPtr.Next = nil
    // 将小于 x 的链表的尾节点指向大于等于 x 的链表的头节点
    lessPtr.Next = greater.Next
    // 返回小于 x 的链表的头节点
    return less.Next
}
```
</details>

<details markdown="1">
<summary>234. 回文链表（Easy）</summary>

> 给你一个单链表的头节点 head (事实上是首元结点)，请你判断该链表是否为回文链表。如果是，返回 true；否则，返回 false

`辅助切片`

```golang
func isPalindrome(head *ListNode) bool {
    vals := []int{}
    for head != nil {
        vals = append(vals, head.Val)
        head = head.Next
    }
    // 遍历切片的前半部分，如果有不相等的值，返回 false
    for i, v := range vals[:len(vals)/2] {
        if v != vals[len(vals)-i-1] {
            return false
        }
    }
    return true
}
```

> 遍历链表，将值存入切片，再判断切片是否为回文

`快慢指针`

```golang
func isPalindrome(head *ListNode) bool {
    var prev *ListNode
    slow, fast := head, head
    // 快指针走到链表尾部，慢指针走到链表中间
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        // 在遍历过程中，将前半部分链表反转
        next := slow.Next
        slow.Next = prev
        prev = slow
        slow = next
    }
    // head 指向前半部分已经反转的首元结点
    head = prev
    // prev 再指回slow，用于再次恢复前半部分的链表
    prev = slow
    // 如果 fast 不为 nil 说明跳出循环的时候 fast.Next 为 nil，即链表长度为奇数
    // 将 slow 跳过中间节点，指向后半部分的首元结点
    if fast != nil {
        slow = slow.Next
    }
    palindrome := true
    // 从中间向两边遍历，判断是否是回文链表
    for head != nil {
        if head.Val != slow.Val {
            palindrome = false
        }
        // 再将前半部分反转回来
        next := head.Next
        head.Next = prev
        prev = head
        head = next // 同时将 head 往后移动（前半部分）
        slow = slow.Next // slow 往后移动（后半部分）
    }
    // 跳出循环，说明是回文链表，返回 true
    return palindrome
}
```

</details>

<details markdown="1">
<summary>141. 环形链表（Easy）</summary>

> 给你一个链表的头节点 head ，判断链表中是否有环。

`快慢指针`

```golang
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

`集合`

```golang
func hasCycle(head *ListNode) bool {
    set := make(map[*ListNode]struct{})
    for head != nil {
        if _, ok := set[head]; ok {
            return true
        }
        set[head] = struct{}{}
        head = head.Next
    }
    return false
}
```

</details>

<details markdown="1">
<summary>142. 环形链表 II（Medium）</summary>

> 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
    
`快慢指针`

```golang
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        // 快慢指针相遇
        if slow == fast {
            // 从 head 和相遇点同时出发
            p := head
            // 再次相遇即为环的入口
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p
        }
    }
    return nil
}
```

`哈希表`
```golang
func detectCycle(head *ListNode) *ListNode {
    seen := map[*ListNode]bool{}
    for head != nil {
        if seen[head] {
            return head
        }
        seen[head] = true
        head = head.Next
    }
    return nil
}
```

</details>

<details markdown="1">
<summary>21. 合并两个有序链表（Easy）</summary>

> 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

```golang
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    dummy := new(ListNode)
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            cur.Next = l1
            cur = cur.Next
            l1 = l1.Next
        } else {
            cur.Next = l2
            cur = cur.Next
            l2 = l2.Next
        }
    }
    // 有一个链表为空时，将另一个链表剩余的值放入新链表中
    switch {
    case l1 != nil:
        cur.Next = l1
    case l2 != nil:
        cur.Next = l2
    }
    return dummy.Next
}
```
</details>

<details markdown="1">
<summary>2. 两数相加（Medium）</summary>

> 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

> 请你将两个数相加，并以相同形式返回一个表示和的链表。你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

```golang
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    dummy := new(ListNode)
    cur := dummy
    var carry int
    for l1 != nil || l2 != nil {
        // 创建新节点用于存放相加后的值，因为返回一个新链表，所以每次都要创建一个新节点
        cur.Next = new(ListNode)
        cur = cur.Next
        if l1 != nil {
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            carry += l2.Val
            l2 = l2.Next
        }
        // 将相加后的值存入新节点
        cur.Val = carry % 10
        // 计算进位值
        carry /= 10
    }
    // 如果最后还有进位值，就再创建一个新节点，此时 carry 不需要再除以 10
    if carry > 0 {
        cur.Next = &ListNode{Val: carry}
    }
    return dummy.Next
}
```
</details>

<details markdown="1">
<summary>19. 删除链表的倒数第N个结点（Medium）</summary>

> 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

```golang
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    for range n {
        head = head.Next
    }
    prev := dummy
    // 当 head 指针指向 nil，此时 prev 指针指向倒数第 n 个节点的前一个节点
    for head != nil {
        head = head.Next
        prev = prev.Next
    }
    // 删除倒数第 n 个节点
    prev.Next = prev.Next.Next
    return dummy.Next
}
```
</details>

<details markdown="1">
<summary>24. 两两交换链表中的结点（Medium）</summary>

> 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

`递归`

```golang
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 最后返回的新链表的头节点
    newHead := head.Next
    // 递归调用，传入的参数是下一次递归的头节点，head 指向下一次递归的头节点
    head.Next = swapPairs(newHead.Next)
    // newHead 指向 head
    newHead.Next = head
    return newHead
}
```

`迭代`

```golang
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    temp := dummy
    // 迭代条件：temp.Next 和 temp.Next.Next 都不为空
    for temp.Next != nil && temp.Next.Next != nil {
        first := temp.Next
        second := temp.Next.Next
        // 交换 first 和 second
        temp.Next = second
        first.Next = second.Next
        second.Next = first
        // first 成为下一轮待交换的前置节点
        temp = first
    }
    return dummy.Next
}
```
</details>

<details markdown="1">
<summary>138. 随机链表的复制（Medium）</summary>

> 给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

```golang
func copyRandomList(head *Node) *Node {
    if head == nil {
        return head
    }
    m := make(map[*Node]*Node)
    curr := head
    // 遍历原链表，将原链表的节点和新链表的节点存储到 map 中，此时新链表的 next 和 random 指针都为空
    for curr != nil {
        node := &Node{Val: curr.Val}
        m[curr] = node
        curr = curr.Next
    }
    curr = head
    // 再次遍历原链表，将新链表的 next 和 random 指针指向正确的节点
    for curr != nil {
        node := m[curr]
        node.Next = m[curr.Next]
        node.Random = m[curr.Random]
        curr = curr.Next
    }
    // 最后直接返回 map 中的头节点，对应的值就是新链表的头节点
    return m[head]
}
```

> 利用 map 存储原链表的节点和新链表的节点，然后再遍历一次，将新链表的 next 和 random 指针指向正确的节点

</details>

<details markdown="1">
<summary>148. 排序链表（Medium）</summary>

> 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

```golang
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    middle := findMiddle(head)
    r := sortList(middle.Next)
    middle.Next = nil
    l := sortList(head)
    return mergeTwoLists(l, r)

}
func findMiddle(head *ListNode) *ListNode {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }
    return slow
}
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
    dummy := new(ListNode)
    curr := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            curr = curr.Next
            l1 = l1.Next
        } else {
            curr.Next = l2
            curr = curr.Next
            l2 = l2.Next
        }
    }
    switch {
    case l1 != nil:
        curr.Next = l1
    case l2 != nil:
        curr.Next = l2
    }
    return dummy.Next
}
```

> 归并排序，找到链表中点，分成两个链表，递归排序，再合并

</details>

<details markdown="1">
<summary>146. LRU缓存（Medium）</summary>

> 设计并实现最近最少使用（LRU）缓存机制。实现 LRUCache 类：

> LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存

> int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1

> void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，

> 则插入该组「关键字-值」。当缓存容量达到上限时，则应该逐出最久未使用的关键字。

> 函数 get(key) 和 put(key, value) 的时间复杂度均为 O(1)

```golang
// 键值对，实现 Value 接口，方便链表操作
type pair struct {
    key, value int
}
// LRU 缓存，使用双向链表和哈希表实现
type LRUCache struct {
    capacity int                   // 缓存容量
    list     *list.List            // 双向链表
    cache    map[int]*list.Element // 哈希表
}
func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity,
        list.New(),
        make(map[int]*list.Element),
    }
}
func (c *LRUCache) Get(key int) int {
    // 如果 key 存在，将节点移到链表头部，并返回节点的值
    if elem, ok := c.cache[key]; ok {
        c.list.MoveToFront(elem)
        // 需要断言，因为 elem.Value 是 interface{} 类型
        return elem.Value.(pair).value
    }
    return -1
}
func (c *LRUCache) Put(key int, value int) {
    // 如果 key 存在，更新节点的值，并将节点移到链表头部
    if elem, ok := c.cache[key]; ok {
        c.list.MoveToFront(elem)
        // 更新对应 key 的 value
        elem.Value = pair{key, value}
        return
    }
    // 添加新节点之前，需要判断缓存是否已满
    // 如果 cache 已满，移除链表尾部节点，并删除哈希表中对应的项以及双向链表中的节点
    // 删除链表尾部节点就符合了 LRU 的要求，因为尾部节点是最久未使用的，每次插入新节点或者更新节点都是在链表头部
    if c.list.Len() == c.capacity {
        last := c.list.Back()
        delete(c.cache, last.Value.(pair).key)
        c.list.Remove(last)
    }
    // 在链表头部插入新节点，并在哈希表中添加 key 和节点的映射
    // PushFront 返回的是 *list.Element，所以可以直接赋值给 map，同时完成了两个操作
    c.cache[key] = c.list.PushFront(pair{key, value})
}

```
</details>

### 二叉树

<details markdown="1">
<summary>94. 二叉树的中序遍历（Easy）</summary>

> 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

`递归`

```golang
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var res []int
    res = append(res, inorderTraversal(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inorderTraversal(root.Right)...)
    return res
}
```

`迭代`

```golang
func inorderTraversal(root *TreeNode) []int {
    var res []int
    var stack []*TreeNode
    for curr := root; curr != nil || len(stack) > 0; {
        // 沿着左子树一直往下走，直到走到叶子节点
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        // 栈顶元素出栈，并访问该节点
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, curr.Val)
        // 访问右子树
        curr = curr.Right
    }
    return res
}
```
</details>

<details markdown="1">
<summary>104. 二叉树的最大深度（Easy）</summary>

> 给定一个二叉树 root ，返回其最大深度。二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

```golang
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
```
</details>

<details markdown="1">
<summary>226. 翻转二叉树（Easy）</summary>

> 给你一颗二叉树的根节点 root，翻转这颗二叉树，并返回其根节点。

```golang
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    // 递归翻转左右子树
    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    return root
}
```
</details>

<details markdown="1">
<summary>101. 对称二叉树（Easy）</summary>

> 给你一个二叉树的根节点 root ，检查它是否轴对称。

```golang
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return symmetric(root.Left, root.Right)
}
func symmetric(p, q *TreeNode) bool {
    switch {
    case p == nil || q == nil:
        return p == q
    case p.Val != q.Val:
        return false
    }
    return symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
}
```
</details>

<details markdown="1">
<summary>543. 二叉树的直径（Easy）</summary>

> 给你一棵二叉树的根节点，返回该树的 直径 。二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。两节点之间路径的 长度 由它们之间边数表示。

```golang
func diameterOfBinaryTree(root *TreeNode) int {
    var res int
    var diameter func(*TreeNode) int
    diameter = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        // 递归计算左右子树的深度
        left := diameter(root.Left)
        right := diameter(root.Right)
        // 计算当前节点的最大直径
        if left+right > res {
            res = left + right
        }
        // 返回当前节点的深度
        depth := left
        if right > depth {
            depth = right
        }
        return depth + 1
    }
    diameter(root)
    return res
}
```
</details>

<details markdown="1">
<summary>102. 二叉树的层序遍历（Medium）</summary>

> 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

```golang
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var res [][]int
    // 模拟队列
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        var level []int
        // 遍历当前层的节点
        for range queue  {
            node := queue[0]
            queue = queue[1:]
            // 将出队元素的值存入 level，记录当前层的值
            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        // 将当前层的值存入 res
        res = append(res, level)
    }
    return res
}
```
</details>

<details markdown="1">
<summary>108. 将有序数组转换为二叉搜索树（Easy）</summary>

> 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

```golang
func sortedArrayToBTS(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) / 2
    return &TreeNode{
        Val:   nums[mid],
        // 递归调用，将数组左边的元素构造左子树，右边的元素构造右子树
        Left:  sortedArrayToBTS(nums[:mid]),
        Right: sortedArrayToBTS(nums[mid+1:]),
    }
}
```
> 中序遍历，总是选择中间位置左边的数字作为根节点

</details>

<details markdown="1">
<summary>129. 求根到叶子节点数字之和（Medium）</summary>

> 给你一个二叉树的根节点 root， 树中每个节点都存放有一个 0 到 9 之间的数字。每条从根节点到叶节点的路径都代表一个数字：例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123。计算从根节点到叶节点生成的 所有数字之和 。

```golang
func sumNumbers(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, prevNum int) int {
    if root == nil {
        return 0
    }
    sum := prevNum*10 + root.Val
    // 如果当前节点是叶子节点，直接返回sum
    if root.Left == nil && root.Right == nil {
        return sum
    }
    // 如果当前节点不是叶子节点，返回左右子树的和，传入的参数是sum作为下次递归的prevNum
    return dfs(root.Left, sum) + dfs(root.Right, sum)
}
```
> DFS，递归计算左右子树的和，传入的参数是当前节点的值

</details>

<details markdown="1">
<summary>98. 验证二叉搜索树（Medium）</summary>

> 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

`递归`

```golang
func isValidBST(root *TreeNode) bool {
    var dfs func(*TreeNode, int, int) bool
    dfs = func(root *TreeNode, min, max int) bool {
        if root == nil {
            return true
        }
        // 如果当前节点的值不在 [min, max] 的范围内，则返回 false
        if root.Val <= min || root.Val >= max {
            return false
        }
        return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
    }
    return dfs(root, -1<<63, 1<<63-1)
}
```

`非递归中序遍历`

```golang
func isValidBST(root *TreeNode) bool {
    var stack []*TreeNode
    var pre *TreeNode
    for len(stack) > 0 || root != nil {
        // 将当前节点的所有左子节点入栈
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 弹出栈顶元素
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 如果当前节点的值小于等于 pre 节点的值，则不是二叉搜索树
        if pre != nil && root.Val <= pre.Val {
            return false
        }
        // 更新 pre 节点
        pre = root
        // 处理右子节点
        root = root.Right
    }
    return true
}
```
</details>

<details markdown="1">
<summary>230. 二叉搜索树中第K小的元素（Medium）</summary>

> 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第k个最小元素（从 1 开始计数）

```golang
func kthSmallest(root *TreeNode, k int) int {
    res := []int{}
    var inorder func(*TreeNode)
    inorder = func(tn *TreeNode) {
        if tn == nil {
            return
        }
        inorder(tn.Left)
        res = append(res, tn.Val)
        inorder(tn.Right)
    }
    inorder(root)
    return res[k-1]
}
```

> 中序遍历，将节点值存入切片，返回第 k 个元素

</details>

<details markdown="1">
<summary>199. 二叉树的右视图（Medium）</summary>

> 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值

```golang
func rightSideView(root *TreeNode) []int {
    var res []int
    var dfs func(*TreeNode, int)
    dfs = func(tn *TreeNode, depth int) {
        if tn == nil {
            return
        }
        if depth == len(res) {
            res = append(res, tn.Val)
        }
        dfs(tn.Right, depth+1)
        dfs(tn.Left, depth+1)
    }
    dfs(root, 0)
    return res
}
```

> 深度优先遍历，每层只取最右边的节点

</details>

<details markdown="1">
<summary>114. 二叉树展开为链表（Medium）</summary>

> 给你二叉树的根结点 root ，请你将它展开为一个单链表：展开后的单链表应该同样使用 TreeNode，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。展开后的单链表应该与二叉树 先序遍历 顺序相同。

```golang
func flatten(root *TreeNode) {
    curr := root
    // 通过右指针遍历
    for curr != nil {
        // 提前记录当前节点的右子树
        right := curr.Right
        // 将当前节点的左子树插入到右子树的地方
        curr.Left, curr.Right = nil, curr.Left
        // 将原来的右子树接到当前右子树的最右边节点
        prev := curr
        for prev.Right != nil {
            prev = prev.Right
        }
        // 将原来的右子树接到当前右子树的最右边节点
        prev.Right = right
        curr = curr.Right
    }
}
```
</details>

<details markdown="1">
<summary>105. 从前序与中序遍历序列构造二叉树（Medium）</summary>

> 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

```golang
func buildTree(preorder, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    // 从中序遍历中找到根节点的索引，根节点即为前序遍历的第一个元素
    i := func(order []int, v int) int {
        var index int
        for order[index] != v {
            index++
        }
        return index
    }(inorder, preorder[0])
    // 根节点为 preorder[0], 中序遍历的根节点索引为 i
    // 前序遍历的左子树为 preorder[1:i+1], 中序遍历的左子树为 inorder[:i]
    // 前序遍历的右子树为 preorder[i+1:], 中序遍历的右子树为 inorder[i+1:]
    return &TreeNode{
        Val:   preorder[0],
        Left:  buildTree(preorder[1:i+1], inorder[:i]),
        Right: buildTree(preorder[i+1:], inorder[i+1:]),
    }
}
```
</details>

<details markdown="1">
<summary>112. 路径总和（Easy）</summary>

> 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

```golang
func hasPathSum(root *TreeNode, sum int) bool {
    // 空树，直接返回false
    if root == nil {
        return false
    }
    // 只存在根节点，判断根节点的值是否等于sum
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    // 递归判断左右子树
    sum -= root.Val
    return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```
</details>

<details markdown="1">
<summary>113. 路径总和 II（Medium）</summary>

> 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。叶子节点 是指没有子节点的节点。

```golang
func pathSum(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return nil
    }
	// 只存在根节点，判断根节点的值是否等于 targetSum
    if root.Left == nil && root.Right == nil {
        if root.Val != targetSum {
            return nil
        }
        return [][]int{{root.Val}}
    }
    var res [][]int
    targetSum -= root.Val
    // 辅助函数，用于处理子树
    process := func(subtree *TreeNode) {
        for _, path := range pathSum(subtree, targetSum) {
            path = append([]int{root.Val}, path...)
            res = append(res, path)
        }
    }
    // 处理左子树和右子树
    process(root.Left)
    process(root.Right)
    return res
}
```

> 递归调用，处理左右子树，将根节点的值加入路径中 

</details>

<details markdown="1">
<summary>437. 路径总和 III（Medium）</summary>

> 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的路径的数目。 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

```golang
func pathSum(root *TreeNode, targetSum int) int {
    preSumMap := map[int]int{0: 1}
    var f func(*TreeNode, int) int
    f = func(root *TreeNode, curSum int) int {
        var ans int
        if root == nil {
            return 0
        }
        curSum += root.Val
        // 这条路径即为 curSum - (curSum - targetSum) = targetSum
        if cnt, ok := preSumMap[curSum-targetSum]; ok {
            ans += cnt
        }
        // 更新前缀和 curSum 的次数
        preSumMap[curSum]++
        // 递归左右子树
        ans += f(root.Left, curSum)
        ans += f(root.Right, curSum)
        // 回溯，恢复状态
        preSumMap[curSum]--
        return ans
    }
    return f(root, 0)
}
```

> 递归，使用哈希表存储前缀和，递归遍历二叉树，计算路径和等于目标值的路径数

</details>

<details markdown="1">
<summary>236. 二叉树的最近公共祖先（Medium）</summary>

> 给定一个二叉树，找到该树中两个指定节点的最近公共祖先。

```golang
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}
```

> 递归，如果左子树和右子树都不为空，说明 p 和 q 分别在左右子树中，返回 root，否则返回不为空的子树

</details>

### 图论

<details markdown="1">
<summary>200. 岛屿数量（Medium）</summary>

>给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

> 可以假设该网格的四条边均被水包围。

```golang
func numIslands(grid [][]byte) int {
    visited := make([][]bool, len(grid))
    for i := range visited {
        visited[i] = make([]bool, len(grid[i]))
    }
    var num int
    for i, r := range grid {
        for j, c := range r {
            if c == '0' || visited[i][j] {
                continue
            }
            num++
            visit(grid, visited, i, j)
        }
    }
    return num
}
func visit(grid [][]byte, visited [][]bool, i, j int) {
    if grid[i][j] == '0' || visited[i][j] {
        return
    }
    visited[i][j] = true
    if i > 0 {
        visit(grid, visited, i-1, j)
    }
    if i < len(grid)-1 {
        visit(grid, visited, i+1, j)
    }
    if j > 0 {
        visit(grid, visited, i, j-1)
    }
    if j < len(grid[i])-1 {
        visit(grid, visited, i, j+1)
    }
}
```
> 深度优先遍历，遍历二维网格，遇到岛屿时，递归遍历相邻的岛屿，标记为已访问

</details>

<details markdown="1">
<summary>207. 课程表（Medium）</summary>

> 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。在选修某些课程之前需要一些先修课程。先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

> 判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 

```golang
func canFinish(numCourses int, prerequisites [][]int) bool {
    edges := make([][]int, numCourses)
    inDegree := make([]int, numCourses)
    for _, info := range prerequisites {
        edges[info[1]] = append(edges[info[1]], info[0])
        inDegree[info[0]]++
    }
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        numCourses--
        for _, next := range edges[node] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return numCourses == 0
}
```

> 基于 BFS 的拓扑排序，统计每个节点的入度，将入度为 0 的节点加入队列，遍历队列，将入度为 0 的节点出队，更新其邻接节点的入度，如果入度为 0，加入队列，最后判断是否所有节点都入队

</details>

<details markdown="1">
<summary>208. 实现 Trie (前缀树)（Medium）</summary>

> 实现 Trie 类，包含 insert、search、startsWith 方法

```golang
type Trie struct {
    child [26]*Trie // 子节点
    isEnd bool      // 是否是单词结尾
}
func Constructor() Trie {
    return Trie{}
}
func (t *Trie) Insert(word string) {
    node := t
    for _, ch := range word {
        ch -= 'a'
        if node.child[ch] == nil {
            node.child[ch] = &Trie{}
        }
        node = node.child[ch]
    }
    // 遍历完，将当前节点标记为字符串的结尾
    node.isEnd = true
}
func (t *Trie) SearchPrefix(prefix string) *Trie {
    node := t
    for _, ch := range prefix {
        ch -= 'a'
        if node.child[ch] == nil {
            return nil
        }
        node = node.child[ch]
    }
    // 返回最后一个节点，即前缀的最后一个字符
    return node
}
func (t *Trie) Search(word string) bool {
    node := t.SearchPrefix(word)
    // node != nil 说明 word 的所有前缀都存在, 且最后一个前缀的 isEnd 为 true
    return node != nil && node.isEnd
}
func (t *Trie) StartsWith(prefix string) bool {
    return t.SearchPrefix(prefix) != nil
}
```
</details>

### 回溯

<details markdown="1">
<summary>46. 全排列（Medium）</summary>

> 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

```golang
func permute(nums []int) [][]int {
    var res [][]int
    var f func([]int, []int)
    f = func(nums, path []int) {
        if len(nums) == 0 {
            res = append(res, path)
            return
        }
        for i, v := range nums {
            // 将 nums[:i] 和 nums[i+1:] 拼接起来，得到一个新的切片，再将 v 添加到末尾
            // 这样就得到了一个新的切片，其中不包含原始切片 nums 中的第 i 个数
            newNums := append(append([]int{}, nums[:i]...), nums[i+1:]...)
            newPath := append(path, v)
            f(newNums, newPath)
        }
    }
    f(nums, []int{})
    return res
}
```

> 回溯，递归遍历数组，将每个元素加入路径，再递归遍历剩余元素

</details>

<details markdown="1">
<summary>78. 子集（Medium）</summary>

> 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

```golang
func subsets(nums []int) [][]int {
    sets := make([][]int, 1, 1<<uint(len(nums)))
    for _, num := range nums {
        for _, set := range sets {
            s := make([]int, len(set), len(set)+1)
            copy(s, set)
            // 首先将 num 添加到 s 中，然后将 s 添加到 sets 中
            sets = append(sets, append(s, num))
        }
    }
    return sets
}
```
</details>

<details markdown="1">
<summary>17. 电话号码的字母组合（Medium）</summary>

> 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
> 给出数字到字母的映射如下（与电话按键相同）。

> 1:!@#   2:abc   3:def

> 4:ghi   5:jkl   6:mno

> 7:pqrs  8:tuv   9:wxyz

> *: +    0: _    #: %

```golang
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return nil
    }
    buttons := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    var results []string
    temp := make([]byte, len(digits))
    var dfs func(int)
    dfs = func(i int) {
        // 如果已经遍历完最后一位，此时 i+1=digits，把结果加入到结果集中
        if i == len(digits) {
            results = append(results, string(temp))
            return
        }
        // 获取当前数字对应字母的 byte 数组，因为数字 9 键的字母是从 2 开始的，所以要减去 2
        letters := buttons[digits[i]-'2']
        for j := 0; j < len(letters); j++ {
            temp[i] = letters[j]
            dfs(i + 1)
        }
    }
    // 从参数 digits 的第一位开始遍历
    dfs(0)
    return results
}
```

> 回溯，递归遍历数字对应的字母，将每个字母加入路径，再递归遍历下一个数字

</details>

<details markdown="1">
<summary>39. 组合总和（Medium）</summary>

> 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。candidates 中的数字可以无限制重复被选取。

```golang
func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    var dfs func([]int, int, int)
    dfs = func(comb []int, index, target int) {
        // 如果 target 为 0，说明找到了一个组合，将它放入结果中，然后返回
        if target == 0 {
            res = append(res, append([]int{}, comb...))
            return
        }
        // 从 index 开始遍历 candidates
        for i, c := range candidates[index:] {
            if c <= target {
                // 注意这里的 index+i，因为 candidates 中的数字可以重复使用，所以下一轮搜索的起点仍然是 index+i
                // target - c 为下一轮搜索的目标
                dfs(append(comb, c), index+i, target-c)
            }
        }
    }
    dfs(nil, 0, target)
    return res
}
```
</details>

<details markdown="1">
<summary>22. 括号生成（Medium）</summary>

> 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

```golang
func generateParenthesis(n int) []string {
    pair := make([]byte, n*2)
    var dfs func([]string, []byte, int, int, int) []string
    dfs = func(pairs []string, pair []byte, n, left, right int) []string {
        // 如果左和右括号都用完了，就加入到结果中
        if left == n && right == n {
            return append(pairs, string(pair))
        }
        // 如果左括号还有剩余，就可以放一个左括号
        if left < n {
            pair[left+right] = '('
            pairs = dfs(pairs, pair, n, left+1, right)
        }
        // 如果右括号的数量小于左括号的数量，就可以放一个右括号
        if right < left {
            pair[left+right] = ')'
            pairs = dfs(pairs, pair, n, left, right+1)
        }
        return pairs
    }
    return dfs(nil, pair, n, 0, 0)
}
```
</details>

<details markdown="1">
<summary>79. 单词搜索（Medium）</summary>

> 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

```golang
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    used := make([][]bool, m)
    for i := range used {
        used[i] = make([]bool, n)
    }
    var canFind func(r, c, i int) bool
    canFind = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }
        if r < 0 || r >= m || c < 0 || c >= n {
            return false
        }
        // 已经访问过或者不符合当前字符
        if used[r][c] || board[r][c] != word[i] {
            return false
        }
        used[r][c] = true
        // 间接实现了回溯和剪枝
        if canFind(r-1, c, i+1) || canFind(r+1, c, i+1) || canFind(r, c-1, i+1) || canFind(r, c+1, i+1) {
            return true
        } else {
            used[r][c] = false // 重新标记为未访问, 因为下一次可能会访问到
            return false
        }
    }
    // 遍历所有的起点
    for i := range board {
        for j := range board[i] {
            if canFind(i, j, 0) {
                return true
            }
        }
    }
    return false
}
```
</details>

<details markdown="1">
<summary>131. 分割回文串（Medium）</summary>

> 给你一个字符串 s, 请你将 s 分割成一些子串, 使每个子串都是回文串. 返回 s 所有可能的分割方案。

```golang
func partition(s string) [][]string {
    path := []string{}
    ans := [][]string{}
    n := len(s)
    var dfs func(int, int)
    dfs = func(index, start int) {
        if index == n {
            ans = append(ans, append([]string(nil), path...))
            return
        }
        // 不选 index 和 index+1 之间的逗号
        if index < n-1 {
            dfs(index+1, start)
        }
        // 选 index 和 index+1 之间的逗号, 把 s[index] 作为子串的最后一个字符
        if isPalindrome(s, start, index) {
            path = append(path, s[start:index+1])
            dfs(index+1, index+1)
            path = path[:len(path)-1]
        }
    }
    dfs(0, 0)
    return ans
}
func isPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```
</details>

### 二分查找

<details markdown="1">
<summary>35. 搜索插入位置（Easy）</summary>

> 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。请必须使用时间复杂度为 O(log n) 的算法。

```golang
func searchInsert(nums []int, target int) int {
    i, j := 0, len(nums)
    for i < j {
        mid := int(uint(i+j) >> 1)
        switch {
        case nums[mid] < target:
            i = mid + 1
        case nums[mid] > target:
            j = mid
        default:
            return mid
        }
    }
    return i
}
```
</details>

<details markdown="1">
<summary>74. 搜索二维矩阵（Medium）</summary>

> 给你一个满足下述两条属性的 m x n 整数矩阵：每行中的整数从左到右按非递减顺序排列。每行的第一个整数大于前一行的最后一个整数。给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

```golang
func searchMatrix(matrix [][]int, target int) bool {
    for _, row := range matrix {
        i := sort.SearchInts(row, target)
        if i < len(row) && row[i] == target {
            return true
        }
    }
    return false
}
```
</details>

<details markdown="1">
<summary>34. 在排序数组中查找元素的第一个和最后一个位置（Medium）</summary>

> 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。如果数组中不存在目标值 target，返回 [-1, -1]。你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

```golang
func searchRange(nums []int, target int) []int {
    // 使用 SearchInts 在 nums 中搜索 target，如果找到则返回其索引，否则返回将会插入的索引
    left := sort.SearchInts(nums, target)
    if left == len(nums) || nums[left] != target {
        return []int{-1, -1}
    }
    // 按照题目意思，因为存在左索引了，所以不需要再对右索引进行判断是否越界和是否等于 target
    // 找到 target+1 的索引，再往前移动一个位置, 即为 target 的最右索引
    right := sort.SearchInts(nums, target+1) - 1
    return []int{left, right}
}
```
</details>

<details markdown="1">
<summary>33. 搜索旋转排序数组（Medium）</summary>

> 整数数组 nums 按升序排列，数组中的值 互不相同 。在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（0 下标开始）。给你 搜索目标值 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

```golang
func search(nums []int, target int) int {
    lo, hi := 0, len(nums)
    for lo < hi {
        mid := int(uint(lo+hi) >> 1)
        if nums[mid] == target {
            return mid
        }
        if nums[0] <= nums[mid] {
            if nums[0] <= target && target < nums[mid] {
                hi = mid
            } else {
                lo = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[len(nums)-1] {
                lo = mid + 1
            } else {
                hi = mid
            }
        }
    }
    return -1
}
```
> 二分查找，根据 nums[0] 和 nums[mid] 的关系判断 target 在左半部分还是右半部分，再根据 target 和 nums[mid] 的关系更新 lo 和 hi

</details>

<details markdown="1">
<summary>153. 寻找旋转排序数组中的最小值（Medium）</summary>

> 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。

```golang
func findMin(nums []int) int {
    lo, hi := 0, len(nums)-1
    for lo < hi {
        mid := int(uint(lo+hi) >> 1)
        // 如果中间值大于最右边的值，说明最小值在右边
        if nums[mid] > nums[hi] {
            lo = mid + 1
        } else {
            // 包含 nums[mid] == nums[hi] 的情况，所以不能用 hi = mid - 1
            hi = mid
        }
    }
    return nums[lo]
}
```
</details>

<details markdown="1">
<summary>162. 寻找峰值（Medium）</summary>

> 峰值元素是指其值大于左右相邻值的元素。返回数组中的任何一个峰值即可

```golang
func findPeakElement(nums []int) int {
    l, r := 0, len(nums)-1
    for l < r {
        h := int(uint(l+r) >> 1)
        if nums[h] > nums[h+1] {
            r = h
        } else {
            l = h + 1
        }
    }
    return l
}
```

> 二分查找，如果中间值大于右边值，说明峰值在左边，否则在右边

</details>

### 栈

<details markdown="1">
<summary>20. 有效的括号（Easy）</summary>

> 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。有效字符串需满足：左括号必须用相同类型的右括号闭合。左括号必须以正确的顺序闭合。每个右括号都有一个对应的相同类型的左括号。

```golang
func isValid(s string) bool {
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := []byte{}
    for _, ch := range []byte(s) {
        if pair, ok := pairs[ch]; ok {
            if len(stack) == 0 || stack[len(stack)-1] != pair {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, ch)
        }
    }
    return len(stack) == 0
}
```

> 使用栈，遍历字符串，遇到左括号入栈，遇到右括号出栈，判断是否匹配

</details>

<details markdown="1">
<summary>155. 最小栈（Medium）</summary>

> 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

```golang
type MinStack struct {
    stack    []int
    minstack []int // 辅助栈，存储最小值
}
func Constructor() MinStack {
    return MinStack{
        stack:    []int{},
        minstack: []int{math.MaxInt64},
    }
}
// 当一个元素要入栈时，首先取当前辅助栈的栈顶存储的最小值，然后与当前元素比较出最小值
// 将得出的最小值推入辅助栈的栈顶，即辅助栈的栈顶存储当前栈所对应的最小值
func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    top := this.minstack[len(this.minstack)-1]
    this.minstack = append(this.minstack, min(val, top))
}
// 删除堆栈顶部的元素，将辅助栈的栈顶元素一同弹出
func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack)-1]
    this.minstack = this.minstack[:len(this.minstack)-1]
}
func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin() int {
    return this.minstack[len(this.minstack)-1]
}
```
</details>

<details markdown="1">
<summary>394. 字符串解码（Medium）</summary>

> 给定一个经过编码的字符串，返回它解码后的字符串。编码规则是: k[encoded_string] ，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。你可以认为输入字符串总是有效的；没有额外的空格，方括号格式正确等。
```golang
func decodeString(s string) string {
    cntStack, strStack := []int{}, []string{}
    currNum, currStr := 0, ""
    for _, v := range s {
        switch {
        case v >= '0' && v <= '9':
            currNum = currNum*10 + int(v-'0')
        case v == '[':
            cntStack = append(cntStack, currNum)
            strStack = append(strStack, currStr)
            currNum, currStr = 0, ""
        case v == ']':
            // 从数字栈中弹出一个重复次数，从字符串栈中弹出一个字符串
            num, str := cntStack[len(cntStack)-1], strStack[len(strStack)-1]
            cntStack, strStack = cntStack[:len(cntStack)-1], strStack[:len(strStack)-1]
            // 将当前的字符串重复指定的次数，然后与上一个字符串合并
            currStr = str + strings.Repeat(currStr, num)
        default:
            currStr += string(v)
        }
    }
    return currStr
}
```

> 使用两个栈，一个存储数字，一个存储字符串，遍历字符串，遇到数字入数字栈，遇到左括号入字符串栈，遇到右括号出栈，重复字符串，最后返回结果

</details>

<details markdown="1">
<summary>739. 每日温度（Medium）</summary>

> 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。 如果气温在这之后都不会升高，请在该位置用 0 来代替。

```golang
func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := []int{}
    for i, v := range temperatures {
        for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
            // 当前温度大于栈顶元素对应的温度时，栈顶元素出栈，计算结果
            res[stack[len(stack)-1]] = i - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}
```

> 单调栈，遍历数组，维护一个单调递减栈，栈中存储的是下标，遇到比栈顶元素大的元素，出栈并计算结果

</details>

### 堆

<details markdown="1">
<summary>215. 数组中的第K个最大元素（Medium）</summary>

> 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。请设计时间复杂度为 O(n) 的算法解决此问题

```golang
func findKthLargest(nums []int, k int) int {
    lo, hi := 0, len(nums)-1
    for lo < hi {
        pivot := partition(nums, lo, hi)
        switch {
        // k-1 是下标，mid 是下标，所以 k-1 == mid 时，找到了第 k 个最大元素
        case k-1 < pivot:
            hi = pivot - 1
        case k-1 > pivot:
            lo = pivot + 1
        default:
            return nums[pivot]
        }
    }
    return nums[lo]
}
// 一次划分，返回枢轴元素的下标，枢轴的位置是确定的，后续不再变化
// 枢轴左边的元素都大于枢轴，右边的元素都小于枢轴，不需要考虑枢轴两侧元素的顺序
func partition(nums []int, lo, hi int) int {
    pivot := nums[lo]
    i, j := lo, hi
    for i < j {
        for i < j && nums[j] <= pivot {
            j--
        }
        nums[i] = nums[j]
        for i < j && nums[i] >= pivot {
            i++
        }
        nums[j] = nums[i]
    }
    nums[i] = pivot
    return i
}
```

> 快速选择，使用快速排序的 partition 函数，根据 pivot 的位置判断 k 在左边还是右边，递归处理
</details>

### 贪心算法

<details markdown="1">
<summary>121. 买卖股票的最佳时机（Easy）</summary>

> 给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。设计一个算法来计算你所能获取的最大利润。返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

```golang
func maxProfit(prices []int) int {
    var minIndex, bonus int
    for i, p := range prices {
        // 首先列出 profit 的计算公式，当前价格减去最小价格
        profit := p - prices[minIndex]
        if profit > bonus {
            bonus = profit
        } else if profit < 0 {
            // 更新最小价格的索引
            minIndex = i
        }
    }
    return bonus
}
```
</details>

<details markdown="1">
<summary>55. 跳跃游戏（Medium）</summary>

> 给定一个非负整数数组 nums ，你最初位于数组的第一个下标。数组中的每个元素代表你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标。

```golang
func canJump(nums []int) bool {
    // 定义一个变量 cover，表示当前能够覆盖的最远位置，index 表示数组的最后一个下标
    cover, index := 0, len(nums)-1
    for i := 0; i <= cover; i++ {
        cover = max(cover, i+nums[i])
        // 如果 cover 大于等于 index，说明可以到达最后一个下标，返回 true
        if cover >= index {
            return true
        }
    }
    return false
}
```
</details>

<details markdown="1">
<summary>45. 跳跃游戏 II（Medium）</summary>

> 给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。你的目标是使用最少的跳跃次数到达数组的最后一个位置。

```golang
func jump(nums []int) int {
    steps, position := 0, len(nums)-1
    for position > 0 {
        for i := 0; i < position; i++ {
            // 如果当前位置能够到达的最远位置大于等于 position，说明从当前位置可以一步跳到最后一个位置
            if i+nums[i] >= position {
                position = i
                steps++
                break
            }
        }
    }
    return steps
}
```
</details>

<details markdown="1">
<summary>763. 划分字母区间（Medium）</summary>

> 给你一个字符串 s，请你对 s 的子串进行划分，并返回子串的个数。划分要尽可能多的子串，并且每个字母只能出现在一个子串中。将划分结果按顺序连接，得到的字符串应当与原字符串完全相同。

```golang
func partitionLabels(s string) []int {
    lastIndex := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        lastIndex[s[i]] = i
    }
    var ans []int
    start, end := 0, 0
    // 遍历字符串，end 记录当前子串的结束位置
    for i := 0; i < len(s); i++ {
        end = max(end, lastIndex[s[i]])
        // 当前位置等于 end 时，表示当前子串结束，可以划分
        if i == end {
            ans = append(ans, end-start+1)
            start = end + 1
        }
    }
    return ans
}
```
> 贪心算法，遍历字符串，记录每个字符最后出现的位置，遍历字符串，更新当前子串的结束位置，当当前位置等于结束位置时，划分子串

</details>

### 动态规划

<details markdown="1">
<summary>70. 爬楼梯（Easy）</summary>

> 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

```golang
func climbStairs(n int) int {
    // p:dp[i-2]，q:dp[i-1]，r:dp[i]
    p, q, r := 0, 0, 1
    for i := 1; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}
```

> 动态规划，dp[i] = dp[i-1] + dp[i-2]

</details>

<details markdown="1">
<summary>118. 杨辉三角（Easy）</summary>

> 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

```golang
func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        ans[i] = make([]int, i+1)
        ans[i][0], ans[i][i] = 1, 1
        for j := 1; j < i; j++ {
            ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
        }
    }
    return ans
}
```
</details>

<details markdown="1">
<summary>198. 打家劫舍（Medium）</summary>

> 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统， 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

```golang
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    // 当房间数大于等于 2
    first, second := nums[0], max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        // first 表示前 i-1 个房间的最大值，second 表示前 i 个房间的最大值
        first, second = second, max(first+nums[i], second)
    }
    return second
}
```
</details>

<details markdown="1">
<summary>279. 完全平方数（Medium）</summary>

> 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。

```golang
func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        cnt := math.MaxInt32
        for j := 1; j*j <= i; j++ {
            cnt = min(cnt, dp[i-j*j])
        }
        dp[i] = cnt + 1
    }
    return dp[n]
}
```
</details>

<details markdown="1">
<summary>322. 零钱兑换（Medium）</summary>

> 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。你可以认为每种硬币的数量是无限的。

```golang
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    dp := make([]int, amount+1)
    for i := 1; i <= amount; i++ {
        // 初始化为 amount+1，不可能取到的值
        dp[i] = amount + 1
    }
    for i := 1; i <= amount; i++ {
        for _, coin := range coins {
            if i >= coin {
                // dp[i-coin]+1：使用一枚 coin 面值的硬币，然后计算剩余金额 i-coin 的最少硬币数
                dp[i] = min(dp[i], dp[i-coin]+1)
            }
        }
    }
    if dp[amount] == amount+1 {
        return -1
    }
    return dp[amount]
}
```
</details>

<details markdown="1">
<summary>139. 单词拆分（Medium）</summary>

> 给你一个字符串 s 和一个字符串列表 wordDict ，s 的所有字符都是小写英文字母。 如果可以利用 wordDict 中的单词拼接 s ，则返回 true ；否则，返回 false 。

```golang
func wordBreak(s string, wordDict []string) bool {
    set := make(map[string]struct{})
    for _, word := range wordDict {
        set[word] = struct{}{}
    }
    // ans[i] 表示 s[:i] 是否可以拆分成 wordDict 中的单词
    ans := make([]bool, len(s)+1)
    ans[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            // 如果 ans[j] 为 true 且 s[j:i] 在 wordDict 中，则 ans[i] 为 true
            if _, ok := set[s[j:i]]; ok && ans[j] {
                ans[i] = true
                break
            }
        }
    }
    return ans[len(s)]
}
```
</details>

<details markdown="1">
<summary>300. 最长递增子序列（Medium）</summary>

> 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

```golang
func lengthOfLIS(nums []int) int {
    var res int
    // dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
    dp := make([]int, len(nums))
    // 初始化 dp 数组，每个元素都至少可以单独成为一个子序列
    for i := range dp {
        dp[i] = 1
    }
    for i := range nums {
        for j := range i {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    for _, v := range dp {
        res = max(res, v)
    }
    return res
}
```
</details>

<details markdown="1">
<summary>416. 分割等和子集（Medium）</summary>

> 给你一个只包含正整数的非空数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

```golang
// 如果数组的和是奇数，那么不可能分割成两个和相等的子集
// 如果数组的和是偶数，那么问题转化成背包问题，背包容量是数组和的一半
func canPartition(nums []int) bool {
    var sum int
    for _, num := range nums {
        sum += num
    }
    if sum%2 == 1 {
        return false
    }
    target := sum >> 1
    // dp 表示是否可以从数组中选取若干个数，使得这些数的和等于 target
    dp := make([]bool, target+1)
    dp[0] = true
    for _, num := range nums {
        for j := target; j >= num; j-- {
            dp[j] = dp[j] || dp[j-num]
        }
    }
    return dp[target]
}
```
</details>

### 多维动态规划

<details markdown="1">
<summary>62. 不同路径（Medium）</summary>

`dp`

```golang
// dp[j] 存储的是到达当前行第 j 列的路径数量。
// dp[j] 的新值等于 dp[j]（上边的路径数量, 同一列）和 dp[j-1]（左边列的路径数量）的和。
func uniquePaths(m, n int) int {
    dp := make([]int, n)
    // 到达第一行的任何位置只有一种走法, 所以初始化为 1
    for i := range dp {
        dp[i] = 1
    }
    // 从第二行开始，每个位置的走法等于左边 (dp[j-1]) 和上边 (dp[j]) 的走法之和
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] += dp[j-1]
        }
    }
    return dp[n-1]
}
```

`排列组合`

```golang
func uniquePaths1(m, n int) int {
    return int(new(big.Int).Binomial(int64(m+n-2), int64(m-1)).Int64())
}
```

> 数学方法，排列组合，从左上角到右下角一共需要走 m+n-2 步，其中 m-1 步向下，n-1 步向右，所以答案是 C(m+n-2, m-1)

</details>

<details markdown="1">
<summary>64. 最小路径和（Medium）</summary>

> 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

```golang
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]int, n)
    dp[0] = grid[0][0]
    // 第一行中除了第一个元素外，每个元素的路径和等于左边的路径和加上当前格子的权重
    for i := 1; i < n; i++ {
        dp[i] = dp[i-1] + grid[0][i]
    }
    // 从第二行开始，每个元素的路径和等于上边元素的路径和 (dp[j]) 和左边元素的路径和 (dp[j-1]) 的较小值加上当前格子的权重
    for i := 1; i < m; i++ {
        dp[0] += grid[i][0]
        for j := 1; j < n; j++ {
            dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
        }
    }
    return dp[n-1]
}
```
> 和第 62 题类似，不过是格子加了权重，求最小路径和
 
</details>

<details markdown="1">
<summary>5. 最长回文子串（Medium）</summary>

> 给你一个字符串 s，找到 s 中最长的回文子串。如果一个字符串从左向右写和从右向左写是一样的，这样的字符串就是回文串。

```golang
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        // 对应奇数长度的回文串
        l1, r1 := expandAroundCenter(s, i, i)
        // 对应偶数长度的回文串
        l2, r2 := expandAroundCenter(s, i, i+1)
        if r1-l1 > end-start {
            start, end = l1, r1
        }
        if r2-l2 > end-start {
            start, end = l2, r2
        }
    }
    return s[start : end+1]
}
func expandAroundCenter(s string, left, right int) (int, int) {
    // 从中心向两边扩展
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    // 返回符合条件的子串的左右索引, 逆操作 left--, right++
    return left + 1, right - 1
}
```
</details>

<details markdown="1">
<summary>1143. 最长公共子序列（Medium）</summary>

> 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。如果不存在公共子序列，返回 0 。

> 子序列: 是由原字符串删除一些(或不删除)字符而不改变剩余字符相对位置形成的新字符串

```golang
func longestCommonSubsequence(text1, text2 string) int {
    m, n := len(text1), len(text2)
    // dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列长度
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                // 如果 text1[i-1] == text2[j-1]，则 text1[i-1] 和 text2[j-1] 必然在最长公共子序列中
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                // 否则，text1[i-1] 和 text2[j-1] 至少有一个不在最长公共子序列中
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[m][n]
}
```
</details>

<details markdown="1">
<summary>72. 编辑距离（Medium）</summary>

> 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。 你可以对一个单词进行如下三种操作： 1. 插入一个字符 2. 删除一个字符 3. 替换一个字符

```golang
func minDistance(word1, word2 string) int {
    m, n := len(word1), len(word2)
    // dp[i][j] 表示 word1[:i] 转换成 word2[:j] 所使用的最少操作数
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            // 如果 word1[i-1] == word2[j-1]，则不需要操作
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                // 否则，取插入、删除、替换操作的最小值
                dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
```
</details>

### 技巧

<details markdown="1">
<summary>136. 只出现一次的数字（Easy）</summary>

> 给你一个非空数组，这个数组中有一个数字只出现了一次，其他的数字都出现了两次。找出这个只出现一次的数字。

`位运算`

```golang
func singleNumber(nums []int) int {
    var single int
    for _, num := range nums {
        // 异或运算，相同的数异或结果为0，0和任何数异或结果为任何数
        single ^= num
    }
    return single
}
```

`集合`

```golang
func singleNumber1(nums []int) int {
    set := make(map[int]struct{})
    for _, v := range nums {
        // 如果字典中存在该元素，则删除，否则添加
        if _, ok := set[v]; ok {
            delete(set, v)
        } else {
            set[v] = struct{}{}
        }
    }
    for k := range set {
        return k
    }
    return -1
}

```
</details>

<details markdown="1">
<summary>169. 多数元素（Easy）</summary>

> 给定一个大小为 n 的数组 nums，返回其中的多数元素。多数元素是指在数组中出现次数大于 n/2 的元素。可以假定数组是非空的，并且给定的数组总是存在多数元素。

`摩尔投票法`

```golang
func majorityElement(nums []int) int {
    // 定义 众数 和 统计数
    var major, cnt int
    for _, num := range nums {
        // 如果计数清零则重新选定当前 num 为 major
        if cnt == 0 {
            major = num
        }
        // 如果当前 num 等于 major 则计数加一，否则减一
        if num == major {
            cnt++
        } else {
            cnt--
        }
    }
    return major
}
```

`哈希表`

```golang
func majorityElement1(nums []int) int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
        // 如果当前数字出现次数大于 n/2，则找到众数
        if m[num] > len(nums)/2 {
            return num
        }
    }
    return -1
}
```
</details>