package wallsgates

import "leetcode-go/template"

// 286. 墙与门
// https://leetcode.cn/problems/walls-and-gates/
// 给一个 m x n 网格 rooms：-1 表示墙，0 表示门，INF(2147483647) 表示空房间。
// 原地把每个空房间填成它到最近门的距离，无法到达门则保持 INF。
//
// 核心思路：所有门同时作为起点做多源 BFS，第一次到达某个空房间时的层数就是最近距离，
// 每个格子只入队一次。

// INF 表示空房间。
const INF = 2147483647

// wallsAndGates 解法一：多源 BFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}
	m, n := len(rooms), len(rooms[0])
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range template.Dirs4 {
			r, c := cur[0]+d[0], cur[1]+d[1]
			// 只有尚未访问的空房间才会被更新（INF 即未访问）
			if r < 0 || r >= m || c < 0 || c >= n || rooms[r][c] != INF {
				continue
			}
			rooms[r][c] = rooms[cur[0]][cur[1]] + 1
			queue = append(queue, [2]int{r, c})
		}
	}
}

// wallsAndGatesPerGate 解法二：对每个门单独 BFS 取最小值（对照用）
// 每个门做一次完整 BFS（template.GridBFS），再用得到的距离更新每个空房间的最小值。
// 时间复杂度 O(k*m*n)，k 为门的数量；空间复杂度 O(m*n)。
func wallsAndGatesPerGate(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}
	m, n := len(rooms), len(rooms[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] != 0 {
				continue
			}
			dist := template.GridBFS(m, n, i, j, func(r, c int) bool {
				return rooms[r][c] != -1
			})
			for r := 0; r < m; r++ {
				for c := 0; c < n; c++ {
					if rooms[r][c] > 0 && dist[r][c] != -1 && dist[r][c] < rooms[r][c] {
						rooms[r][c] = dist[r][c]
					}
				}
			}
		}
	}
}
