package openlock

// 752. 打开转盘锁
// https://leetcode.cn/problems/open-the-lock/
// 一个 4 位转盘锁，每位 '0'-'9' 可循环拨动，初始为 "0000"。每次可以把某一位上下拨动一格。
// deadends 中的状态一旦到达锁就会卡死。返回到达 target 的最少拨动次数，无法到达返回 -1。
//
// 核心思路：10^4 个状态，每个状态 8 个邻居，求最短路 → BFS。
// 起点和终点都已知，双向 BFS 每次扩展较小的一侧，可大幅减少搜索的状态数。

// neighbors 返回状态 s 的 8 个相邻状态。
func neighbors(s string) []string {
	res := make([]string, 0, 8)
	bs := []byte(s)
	for i := range bs {
		orig := bs[i]
		for _, delta := range []byte{1, 9} { // +1 和 -1（mod 10）
			bs[i] = '0' + (orig-'0'+delta)%10
			res = append(res, string(bs))
		}
		bs[i] = orig
	}
	return res
}

// openLock 解法一：双向 BFS ⭐ 最优方案
// 时间复杂度 O(10^4*4*2)，空间复杂度 O(10^4)，实际访问状态数远少于单向 BFS。
func openLock(deadends []string, target string) int {
	visited := map[string]bool{}
	for _, d := range deadends {
		visited[d] = true
	}
	if visited["0000"] || visited[target] {
		return -1
	}
	if target == "0000" {
		return 0
	}
	front := map[string]bool{"0000": true}
	back := map[string]bool{target: true}
	visited["0000"], visited[target] = true, true
	for step := 1; len(front) > 0 && len(back) > 0; step++ {
		if len(front) > len(back) {
			front, back = back, front
		}
		next := map[string]bool{}
		for s := range front {
			for _, nb := range neighbors(s) {
				if back[nb] {
					return step
				}
				if !visited[nb] {
					visited[nb] = true
					next[nb] = true
				}
			}
		}
		front = next
	}
	return -1
}

// openLockBFS 解法二：单向 BFS（对照用）
// 从 "0000" 逐层扩展直到遇到 target。
// 时间复杂度 O(10^4*4*2)，空间复杂度 O(10^4)。
func openLockBFS(deadends []string, target string) int {
	visited := map[string]bool{}
	for _, d := range deadends {
		visited[d] = true
	}
	if visited["0000"] {
		return -1
	}
	visited["0000"] = true
	queue := []string{"0000"}
	for step := 0; len(queue) > 0; step++ {
		var next []string
		for _, s := range queue {
			if s == target {
				return step
			}
			for _, nb := range neighbors(s) {
				if !visited[nb] {
					visited[nb] = true
					next = append(next, nb)
				}
			}
		}
		queue = next
	}
	return -1
}
