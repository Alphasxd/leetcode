// 134.加油站 https://leetcode.cn/problems/gas-station/

// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
// 你从其中的一个加油站出发，开始时油箱为空。
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。
// 如果存在解，则 保证 它是 唯一 的。

package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
    left := 0                   // 油箱剩余油量
    start := 0                  // 起始加油站
    totalGas, totalCost := 0, 0 // 总油量，总消耗
    for i := 0; i < len(gas); i++ {
        totalGas += gas[i]
        totalCost += cost[i]
        left += gas[i] - cost[i] // 累加每次剩余的油量
        if left < 0 {            // 如果剩余油量小于0，说明从起始加油站到当前加油站都不能作为起始加油站
            start = i + 1 // 从下一个加油站开始
            left = 0      // 油箱剩余油量清零
        }
    }
    if totalGas < totalCost { // 总油量小于总消耗，说明无法环绕一周
        return -1
    }
    return start // 总油量大于总消耗，说明可以环绕一周，返回起始加油站
}
