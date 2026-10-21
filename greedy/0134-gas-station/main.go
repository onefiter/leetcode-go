package gasstation

// 134. 加油站
// https://leetcode.cn/problems/gas-station/
// 环路上有 n 个加油站，gas[i] 为第 i 站可加的油量，cost[i] 为从 i 开到 i+1 的耗油量。
// 油箱容量无限、初始为空，若能从某站出发绕行一周返回则返回该站下标，否则返回 -1（答案唯一）。
//
// 核心思路：若总油量 sum(gas) < 总消耗 sum(cost)，一定无解；否则一定有解。
// 从 start 出发累计剩余油量 tank，若到 i 时 tank < 0，说明 [start, i] 中任何站出发都到不了 i+1
// （从中间某站 k 出发相当于丢掉了 [start, k) 这段非负的累计油量，只会更少），
// 于是直接把起点跳到 i+1 并清空 tank。总量非负保证最后剩下的起点能绕完一圈。

// canCompleteCircuit 解法一：一次遍历贪心 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func canCompleteCircuit(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0
	for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff
		if tank < 0 { // [start, i] 都不能作为起点
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

// canCompleteCircuitBrute 解法二：枚举起点模拟（对照用）
// 对每个起点模拟绕行一圈，途中油量为负即失败。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func canCompleteCircuitBrute(gas []int, cost []int) int {
	n := len(gas)
	for s := 0; s < n; s++ {
		tank, ok := 0, true
		for k := 0; k < n; k++ {
			i := (s + k) % n
			tank += gas[i] - cost[i]
			if tank < 0 {
				ok = false
				break
			}
		}
		if ok {
			return s
		}
	}
	return -1
}
