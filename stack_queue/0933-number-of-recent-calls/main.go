package recentcounter

// 933. 最近的请求次数
// https://leetcode.cn/problems/number-of-recent-calls/
// 每次 Ping(t) 表示在时刻 t 来了一个请求，返回 [t-3000, t] 内的请求数（含本次）。
// 保证每次 t 严格递增。
//
// 核心思路：t 递增，所以过期的请求一定在最早的那一端——用队列：
// 新请求入队尾，把队头 < t-3000 的请求全部弹出，剩下的个数就是答案。

// RecentCounter 解法一：队列 ⭐ 最优方案
// 每个请求最多入队、出队各一次。
// 时间复杂度 均摊 O(1)（每次 Ping），空间复杂度 O(W)，W 为 3000 毫秒窗口内的请求数。
type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for this.q[0] < t-3000 {
		this.q = this.q[1:]
	}
	return len(this.q)
}
