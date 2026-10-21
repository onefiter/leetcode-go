package grumpyowner

// 1052. 爱生气的书店老板
// https://leetcode.cn/problems/grumpy-bookstore-owner/
// 书店营业 n 分钟，第 i 分钟进店 customers[i] 位顾客；grumpy[i]=1 表示老板这一分钟生气，
// 生气时顾客不满意，不生气时顾客满意。老板有一个秘密技巧，可以连续 minutes 分钟不生气（只能用一次）。
// 返回最多能让多少顾客满意。
// 例如 customers=[1,0,1,2,1,1,7,5], grumpy=[0,1,0,1,0,1,0,1], minutes=3 -> 16。
//
// 一句话理解题意：不生气时的顾客无论如何都满意；技巧只是「挽回」某段长度为 minutes 的窗口里
// 生气时段的顾客，所以要找挽回人数最多的那个窗口。
//
// 核心思路：答案 = 基础满意人数 + 定长窗口内「生气时段顾客数」的最大值。

// maxSatisfied 解法一：定长滑动窗口 ⭐ 最优方案
// base 累加所有不生气分钟的顾客；extra 是当前窗口内生气分钟的顾客数，
// 窗口右移时加上进入的、减去离开的（只计生气分钟），记录 extra 最大值。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	base, extra, best := 0, 0, 0
	for i, c := range customers {
		if grumpy[i] == 0 {
			base += c
		} else {
			extra += c
		}
		// 窗口为 [i-minutes+1, i]，把刚离开窗口的 i-minutes 移出
		if j := i - minutes; j >= 0 && grumpy[j] == 1 {
			extra -= customers[j]
		}
		best = max(best, extra)
	}
	return base + best
}

// maxSatisfiedEnumeration 解法二：枚举技巧的起始分钟（对照用）
// 对每个起点 s，假设 [s, s+minutes) 不生气，重新统计全天满意人数。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func maxSatisfiedEnumeration(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	best := 0
	for s := 0; s == 0 || s+minutes <= n; s++ {
		total := 0
		for i, c := range customers {
			if grumpy[i] == 0 || (i >= s && i < s+minutes) {
				total += c
			}
		}
		best = max(best, total)
	}
	return best
}
