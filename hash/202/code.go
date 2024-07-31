// 202. 快乐数 https://leetcode.cn/problems/happy-number/

// 判断给定的数是否为快乐数。
// 快乐数的定义为：
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果 可以变为  1，那么这个数就是快乐数。

// 示例 1：
// 输入：19
// 输出：true
// 19 -> 82 -> 68 -> 100 -> 1

// 示例 2：
// 输入：n = 2
// 输出：false
// 2 -> 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4

package leetcode

// 哈希表，记录每次计算的结果，如果出现重复的结果，说明不是快乐数
func isHappy0(n int) bool {
    m := map[int]bool{}
    for n != 1 {
        // 如果出现循环，说明不是快乐数
        // 譬如 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4
        if m[n] {
            return false
        }
        m[n] = true
        n = func(n int) int {
            var next int
            for n > 0 {
                t := n % 10
                next += t * t
                n /= 10
            }
            return next
        }(n)
    }
    return true
}

func getNext(n int) int {
    var next int
    for n > 0 {
        t := n % 10
        next += t * t
        n /= 10
    }
    return next
}

// 双指针，记录快慢指针的值，如果相等，说明不是快乐数
func isHappy1(n int) bool {
    slow, fast := n, getNext(n)
    for fast != 1 && slow != fast {
        slow = getNext(slow)
        fast = getNext(getNext(fast))
    }
    return fast == 1
}

// 数学套路，如果不是快乐数，最终会进入 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4 的循环
func isHappy(n int) bool {
    squares := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
    for n != 1 && n != 4 {
        var next int
        for n > 0 {
            t := n % 10
            next += squares[t]
            n /= 10
        }
        n = next
    }
    return n == 1
}
